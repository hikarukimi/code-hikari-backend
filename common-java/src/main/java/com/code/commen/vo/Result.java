package com.code.commen.vo;

import lombok.Data;

@Data
public class Result<T> {
    private T data;
    private int code;
    private String msg;

    public static Result success(Object data) {
        Result result = new Result();
        result.setCode(200);
        result.setMsg("success");
        result.setData(data);
        return result;
    }
    public static Result built(Object data, int code , String msg) {
        Result result = new Result();
        result.setCode(code);
        result.setMsg(msg);
        result.setData(data);
        return result;
    }
}
