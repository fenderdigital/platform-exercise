package com.user.application.service;

import com.user.application.domain.dto.UserDTO;
import com.user.application.models.LoginRequest;
import com.user.application.util.ResponseMessage;

public interface UserService {
    ResponseMessage addNewUser(UserDTO userDTO);

    ResponseMessage updateUser(UserDTO userDTO, String token);

    ResponseMessage deleteUser(String token);

    ResponseMessage login(LoginRequest loginRequest);

    ResponseMessage logout(String token);
}
