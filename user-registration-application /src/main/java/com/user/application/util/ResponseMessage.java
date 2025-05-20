package com.user.application.util;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ResponseMessage {
    private String code;
    private String description;
    private BaseMessage response;
}
