package com.code.ai.controller;

import com.code.ai.service.AiService;
import com.code.commen.dto.AiDto;
import com.code.commen.pojo.AiChatHistories;
import com.code.commen.vo.Result;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("ai")
public class AiController {
    @Autowired
    private AiService aiService;
//  ai提问
    @PostMapping("/chat")
    public Result Chat(@RequestBody AiDto aiDto) {
        String str = aiService.chat(aiDto);
        return Result.success(str);
    }
//  根据用户id查询主题
    @GetMapping("/selectTitle/{id}")
    public Result selectTitle() {
        List<String> data =   aiService.selectTitle();
        return Result.success(data);
    }
//    根据id和主题查找
    @GetMapping("/selectByIdAndTitle/{id}/{title}")
    public Result selectByIdAndTitle(@PathVariable("id") Long id,@PathVariable("title") String title) {
        List<AiChatHistories> data = aiService.selectByIdAndTitle(title);
        return Result.success(data);
    }
}
