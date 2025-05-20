package com.user.application.util;

import com.user.application.domain.UserReg;
import com.user.application.domain.dao.UserDAO;
import com.user.application.domain.dto.UserDTO;

public class Mapper {

    public static UserReg convertUserDTOtoUser(UserDTO userDTO) {
        UserReg user = new UserReg();
        user.setName(userDTO.getName());
        user.setEmail(userDTO.getEmail());
        user.setUsername(userDTO.getUsername());
        user.setPassword(userDTO.getPassword());
        return user;
    }

    public static UserReg updateUser(UserDTO userDTO, UserReg user) {
        user.setName(userDTO.getName());
        user.setEmail(userDTO.getEmail());
        user.setUsername(userDTO.getUsername());
        return user;
    }

    public static UserDAO convertUsertoUserDAO(UserReg user) {
        UserDAO userDAO = new UserDAO();
        userDAO.setId(user.getUserId());
        userDAO.setName(user.getName());
        userDAO.setEmail(user.getEmail());
        userDAO.setUsername(user.getUsername());
        userDAO.setPassword(user.getPassword());
        return userDAO;
    }

}
