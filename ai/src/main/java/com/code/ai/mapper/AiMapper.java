package com.code.ai.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.code.commen.pojo.AiChatHistories;
import org.apache.ibatis.annotations.Mapper;


@Mapper
public interface AiMapper extends BaseMapper<AiChatHistories> {
}
