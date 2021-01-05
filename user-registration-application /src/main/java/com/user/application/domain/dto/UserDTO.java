package com.user.application.domain.dto;

import lombok.Data;

@Data
public class UserDTO {
    private Integer id;
    private String name;
    private String username;
    private String email;
    private String password;
}
