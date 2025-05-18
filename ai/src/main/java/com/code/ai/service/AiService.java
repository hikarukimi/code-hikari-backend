package com.code.ai.service;


import com.code.commen.dto.AiDto;
import com.code.commen.pojo.AiChatHistories;

import java.util.List;

public interface AiService{
    String chat(AiDto aiDto);

    List<String> selectTitle();

    List<AiChatHistories> selectByIdAndTitle(String title);
}
