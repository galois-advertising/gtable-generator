// solopointer1202@gmail.com

namespace galois::gtable {

template <class THandle, class TMessenger>
seek_status_t TableSeeker<THandle, TMessenger>::seek(const THandle& handle, TMessenger* msgr, TQueryResponse* qry_rsp) {
    if (msgr == NULL || qry_rsp == NULL) {
        FATAL("msgr or qry_rsp is NULL!", "");
        return SEEK_ERROR;
    }

    seek_status_t ret = SEEK_SUCC;
    int idx = 0;
    do {
        idx++;
        // 只有SEEK_SUCC\SEEK_LEFT_JOINED_FAIL状态才定义为成功
        ret = do_seek(handle, msgr);
        if (ret != SEEK_SUCC && ret != SEEK_LEFT_JOINED_FAIL) {
            count_seek_failed(msgr);
            AQL_DEBUG_COLLECT_NODE(true, "seek_fail");
            AQL_WRITE_LOG_DEBUG("%s seek self failed!", name());
            continue;
        }

        uint32_t each_qry_rsps_size = 0;
        do {
            bool is_selected = call_conditioners(handle, msgr);
            if (!is_selected) {
                // monitor filter false
                do_monitor(msgr);
                AQL_DEBUG_COLLECT_NODE(true, "do_cond_fail");
                AQL_WRITE_LOG_DEBUG("%s calls conditioners failed!", name());
                ret = SEEK_CONDITION_FAIL;
                continue;
            }

            TQueryResponse* pass_qry_rsp = qry_rsp;
            TQueryResponseList* inner_qry_rsps = 
                qry_rsp->mutable_inner_query_responses();

            if (this->type() == INDEXTABLE_SEEKER 
                    || this->has_enumerate_primary_key()) {
                // check if hit the LIMIT
                if ((uint32_t)inner_qry_rsps->size() >= this->result_limit_num()) {
                    AQL_WRITE_LOG_DEBUG("%s hits result limit[%u]", 
                            name(), this->result_limit_num());
                    do_recount_hit_result_limit(msgr);
                    break;
                }
                pass_qry_rsp = add_one_response_to(inner_qry_rsps);
            }
            
            if (FLAGS_adtable_aql_key_status_monitor && 
                    FLAGS_adtable_aql_key_status_monitor_index && 
                    this->is_table_seeker_starter()) {
                typename std::vector<TTableSeeker*>::iterator seeker_iter = 
                        _joined_table_seekers.begin();
                if (seeker_iter == _joined_table_seekers.end()) {
                    // will Never happen
                    msgr->p_monitor_info->key_status_ids.push_back({0, 0});
                } else {
                    TTableSeeker* seeker = (*seeker_iter); //第一个join表，是被join的主正排表
                    uint64_t pk_id = seeker->get_primary_key_id(handle, msgr);
                    msgr->p_monitor_info->key_status_ids.push_back({idx - 1, pk_id});
                }
            }

            ret = call_joined_table_seekers(handle, msgr, pass_qry_rsp);

            pass_qry_rsp->set_seek_key_idx(idx - 1);// AQL in 操作查询是返回结果附带req偏移量
            if (ret != SEEK_SUCC) {
                if (this->type() == INDEXTABLE_SEEKER
                        || this->has_enumerate_primary_key()) {
                    remove_last_response_from(inner_qry_rsps);
                }
                ret = SEEK_JOIN_FAIL;
                continue;
            }
            
            // 只有SEEK_SUCC才能fill_result
            int filler_ret = call_result_fillers(msgr, pass_qry_rsp);
            if (filler_ret != 0) {
                if (this->is_left_joined()) {
                    continue;
                }

                if (this->type() == INDEXTABLE_SEEKER
                        || this->has_enumerate_primary_key()) {
                    remove_last_response_from(inner_qry_rsps);
                } else {
                    pass_qry_rsp->Clear();
                }

                AQL_DEBUG_COLLECT_NODE(true, "fill_result_fail");
                ret = SEEK_FILL_FAIL;
                if (FLAGS_adtable_aql_key_status_monitor && this->is_table_seeker_starter()) {
                    msgr->p_monitor_info->key_status.push_back(KEY_STATUS_ERROR);
                }
                continue;
            }

            if (FLAGS_adtable_aql_key_status_monitor && this->is_table_seeker_starter()) {
                msgr->p_monitor_info->key_status.push_back(KEY_STATUS_SUCCESS);
            }
            if (FLAGS_adtable_aql_res_tag_monitor &&
                    this->is_table_seeker_starter() &&
                    !msgr->p_monitor_info->monitor_tag.empty()) {
                msgr->p_monitor_info->res_tag_succ.push_back(msgr->_tag_succ);
            }

            AQL_DEBUG_COLLECT_NODE(this->is_table_seeker_starter(), "seek_succ");

            ++each_qry_rsps_size;

            if (this->is_table_seeker_starter()) {
                ++msgr->p_monitor_info->_succ_item_num;
            }

            if (this->type() == INDEXTABLE_SEEKER 
                    || this->has_enumerate_primary_key()) {
                if (each_qry_rsps_size >= this->each_result_limit_num()) {
                    AQL_WRITE_LOG_DEBUG("%s hits each result limit[%u]", 
                            name(), this->each_result_limit_num());
                    do_recount_hit_each_result_limit(msgr);
                    break;
                }
            }
        } while (next_item(handle, msgr)); // next_item in one key
    } while (next_primary_key(handle, msgr)); // seek next_key 

    return ret;
}

template <class THandle, class TMessenger>
bool TableSeeker<THandle, TMessenger>::call_conditioners(const THandle& handle, TMessenger* msgr) {
    if (_conditioner != NULL) {
        return _conditioner->do_condition(handle, msgr);
    } else {
        AQL_WRITE_LOG_DEBUG("%s's conditioner is NULL!", name());
        return true;
    }
}

template <class THandle, class TMessenger>
seek_status_t TableSeeker<THandle, TMessenger>::call_joined_table_seekers(const THandle& handle,TMessenger* msgr, TQueryResponse* qry_rsp) {
    typename std::vector<TTableSeeker*>::iterator seeker_iter = _joined_table_seekers.begin();
    for (; seeker_iter != _joined_table_seekers.end(); ++seeker_iter) {
        TTableSeeker* seeker = (*seeker_iter);
        seek_status_t ret = seeker->seek(handle, msgr, qry_rsp);
        // left join 失败继续前行
        if (ret == SEEK_LEFT_JOINED_FAIL) {
            AQL_WRITE_LOG_DEBUG("%s left joins table %s failed!",
                    name(), seeker->name());
            continue;
        }
        if (ret != SEEK_SUCC) {
            AQL_WRITE_LOG_DEBUG("%s joins table %s failed!",
                    name(), seeker->name());
            return ret;
        }
    }

    return SEEK_SUCC;
}

template <class THandle, class TMessenger>
int TableSeeker<THandle, TMessenger>::call_result_fillers(TMessenger* msgr, TQueryResponse* qry_rsp) const {
    typename std::vector<TResultFiller*>::const_iterator result_filler_iter = _result_fillers.begin();
    int ret = 0;
    for (; result_filler_iter != _result_fillers.end(); ++result_filler_iter) {
        const TResultFiller* filler = (*result_filler_iter);
        ret = (*result_filler_iter)->fill_result(msgr, qry_rsp);
        if (ret != 0) {
            AQL_WRITE_LOG_DEBUG("table %s's filler %s failed!",
                    name(), filler->name());
            return ret;
        }
    }

    return 0;
}

} 
