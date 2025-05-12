package com.code.ai.service.impl;

import com.code.ai.service.AiService;
import com.code.commen.dto.AiDto;
import org.springframework.ai.chat.prompt.Prompt;
import org.springframework.ai.ollama.OllamaChatModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;


@Service
public class AiServiceImpl implements AiService {
    @Autowired
    private OllamaChatModel ollamaChatModel;
    @Override
    public String chat(AiDto aiDto) {
        Prompt prompt = new Prompt(aiDto.getMsg());
        return ollamaChatModel.call(prompt).getResult().getOutput().getContent();
    }
}
