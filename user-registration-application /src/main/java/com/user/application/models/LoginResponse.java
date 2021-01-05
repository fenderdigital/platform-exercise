package com.user.application.models;

import com.user.application.util.BaseMessage;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class LoginResponse extends BaseMessage {
    private String jwt;
}
