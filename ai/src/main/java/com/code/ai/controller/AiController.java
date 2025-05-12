package com.code.ai.controller;

import com.code.ai.service.AiService;
import com.code.commen.dto.AiDto;
import com.code.commen.vo.Result;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("ai")
public class AiController {
    @Autowired
    private AiService aiService;
    @PostMapping("/chat")
    public Result Chat(@RequestBody AiDto aiDto) {
        String str = aiService.chat(aiDto);
        return Result.success(str);
    }
}
