package com.user.application.service.impl;

import com.user.application.cache.TokenHashMap;
import com.user.application.domain.UserReg;
import com.user.application.domain.dao.UserDAO;
import com.user.application.domain.dto.UserDTO;
import com.user.application.models.LoginRequest;
import com.user.application.models.LoginResponse;
import com.user.application.password.PasswordValidationUtility;
import com.user.application.repository.UserRepository;
import com.user.application.service.UserService;
import com.user.application.util.JwtUtility;
import com.user.application.util.Mapper;
import com.user.application.util.ResponseMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Objects;
import java.util.Optional;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private PasswordValidationUtility passwordValidationUtility;

    @Autowired
    private JwtUtility jwtUtility;

    @Autowired
    private TokenHashMap<Integer, String> tokenMap;

    @Override
    public ResponseMessage addNewUser(UserDTO userDTO) {
        ResponseMessage responseMessage;
        UserDAO userDAO;
        UserReg user;
        try {
            Optional<UserReg> existingUser = userRepository.findByEmail(userDTO.getEmail());
            if (existingUser.isPresent()) {
                responseMessage = new ResponseMessage("ERROR", "user already exist", null);
            } else {

                UserReg userToSave = Mapper.convertUserDTOtoUser(userDTO);
                userToSave.setPassword(passwordValidationUtility.hashPassword(userDTO.getPassword(), 2));
                user = userRepository.save(userToSave);
                if (user.getUserId() != null) {
                    userDAO = Mapper.convertUsertoUserDAO(user);
                    responseMessage = new ResponseMessage("SUCCESS", "user added successfully", userDAO);
                } else {
                    userDAO = new UserDAO();
                    responseMessage = new ResponseMessage("ERROR", "user saving error", userDAO);
                }
            }
        } catch (Exception ex) {
            ex.printStackTrace();
            responseMessage = new ResponseMessage("ERROR", ex.getMessage(), null);
        }
        return responseMessage;
    }

    @Override
    public ResponseMessage updateUser(UserDTO userDTO, String token) {
        ResponseMessage responseMessage;
        UserDAO userDAO;
        UserReg user;
        String jwt = jwtUtility.getJwt(token);
        Integer userId = Integer.valueOf(jwtUtility.getUserId(jwt));

        try {
            Optional<UserReg> userOptional = userRepository.findById(userId);
            if (userOptional.isPresent()) {
                user = userRepository.save(Mapper.updateUser(userDTO, userOptional.get()));
                if (user.getUserId() != null) {
                    userDAO = Mapper.convertUsertoUserDAO(user);
                    responseMessage = new ResponseMessage("SUCCESS", "user updated successfully", userDAO);
                } else {
                    responseMessage = new ResponseMessage("ERROR", "user updating error", null);
                }
            } else {
                responseMessage = new ResponseMessage("ERROR", "user doesn't exist", null);
            }

        } catch (Exception ex) {
            ex.printStackTrace();
            responseMessage = new ResponseMessage("ERROR", ex.getMessage(), null);
        }
        return responseMessage;
    }

    @Override
    public ResponseMessage deleteUser(String token) {
        ResponseMessage responseMessage;
        String jwt = jwtUtility.getJwt(token);
        Integer userId = Integer.valueOf(jwtUtility.getUserId(jwt));
        try {
            userRepository.deleteById(userId);
            tokenMap.remove(Integer.valueOf(userId));
            responseMessage = new ResponseMessage("SUCCESS", "user deleted successfully", null);
        } catch (Exception ex) {
            ex.printStackTrace();
            responseMessage = new ResponseMessage("ERROR", ex.getMessage(), null);
        }
        return responseMessage;
    }

    @Override
    public ResponseMessage login(LoginRequest loginRequest) {
        ResponseMessage responseMessage;
        Optional<UserReg> userRegOptional;

        String username = getUserName(loginRequest);
        String email = getEmail(loginRequest);

        userRegOptional = Objects.nonNull(username) ? userRepository.findByUsername(username) :
                Objects.nonNull(email) ? userRepository.findByEmail(email) : Optional.empty();

        if (userRegOptional.isPresent()) {
            UserReg user = userRegOptional.get();
            if (passwordValidationUtility.isPasswordValid(user, loginRequest.getPassword())) {
                String jwt = jwtUtility.generateToken(user);
                try {
                    tokenMap.put(user.getUserId(), jwt);
                    responseMessage = new ResponseMessage("SUCCESS", "user successfully logged in", new LoginResponse(jwt));
                } catch (Exception ex) {
                    ex.printStackTrace();
                    responseMessage = new ResponseMessage("ERROR", ex.getMessage(), null);
                }
            } else {
                responseMessage = new ResponseMessage("ERROR", "user credentials invalid", null);
            }

        } else {
            responseMessage = new ResponseMessage("ERROR", "user doesn't exist", null);
        }

        return responseMessage;
    }

    @Override
    public ResponseMessage logout(String token) {
        String jwt = jwtUtility.getJwt(token);
        String userId = jwtUtility.getUserId(jwt);
        if (Objects.isNull(userId)) {
            return new ResponseMessage("ERROR", "user doesn't exist", null);
        }
        tokenMap.remove(Integer.valueOf(userId));

        return new ResponseMessage("SUCCESS", "user successfully logged out", null);
    }

    private String getEmail(LoginRequest loginRequest) {
        return Optional.of(loginRequest)
                .map(LoginRequest::getEmail)
                .orElse(null);
    }

    private String getUserName(LoginRequest loginRequest) {
        return Optional.of(loginRequest)
                .map(LoginRequest::getUsername)
                .orElse(null);
    }
}
