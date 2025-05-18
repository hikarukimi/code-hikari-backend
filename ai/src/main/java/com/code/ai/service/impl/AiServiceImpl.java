package com.code.ai.service.impl;


import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.code.ai.mapper.AiMapper;
import com.code.ai.service.AiService;
import com.code.commen.dto.AiDto;
import com.code.commen.pojo.AiChatHistories;
import org.springframework.ai.chat.prompt.Prompt;
import org.springframework.ai.ollama.OllamaChatModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;


@Service
public class AiServiceImpl implements AiService {
    @Autowired
    private OllamaChatModel ollamaChatModel;
    @Autowired
    private AiMapper aiMapper;

    @Override
    public String chat(AiDto aiDto) {
        //访问模型得出答案
        Prompt prompt = new Prompt(aiDto.getMsg());
        String answer = ollamaChatModel.call(prompt).getResult().getOutput().getContent();
        //封装存入数据库
        AiChatHistories chat = new AiChatHistories();
        chat.setAnswer(answer);
        chat.setProblem(aiDto.getMsg());
        chat.setTitle(aiDto.getMsg());
        chat.setUserId(aiDto.getUserId());
        aiMapper.insert(chat);
        //返回数据
        return answer;
    }

    @Override
    public List<String> selectTitle() {
        // DOTO 重数据库中查询用户信息
        QueryWrapper<AiChatHistories> wrapper = new QueryWrapper();
        wrapper.eq("user_id", 1)
                .select("title");
        List<AiChatHistories> aiChatHistories = aiMapper.selectList(wrapper);
        List<String> list = new ArrayList<>();
        for (AiChatHistories item : aiChatHistories){
            list.add(item.getTitle());
        }
        return list;
    }

    @Override
    public List<AiChatHistories> selectByIdAndTitle(String title) {
        LambdaQueryWrapper<AiChatHistories> wrapper = new LambdaQueryWrapper();
        wrapper.eq(AiChatHistories::getUserId, 1)
                .eq(AiChatHistories::getTitle, title);
        List<AiChatHistories> data = aiMapper.selectList(wrapper);
        return data;
    }

}
