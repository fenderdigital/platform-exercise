package com.user.application.domain.dao;

import com.user.application.util.BaseMessage;
import lombok.Data;

@Data
public class UserDAO extends BaseMessage {
    private Integer id;
    private String name;
    private String username;
    private String email;
    private String password;
}
