package com.user.application.controller;

import com.user.application.domain.dto.UserDTO;
import com.user.application.models.LoginRequest;
import com.user.application.service.JwtValidationService;
import com.user.application.service.UserService;
import com.user.application.util.ResponseMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Objects;

@RestController
@RequestMapping(path = "/user")
public class UserController {

    @Autowired
    private UserService userService;

    @Autowired
    private JwtValidationService jwtValidationService;


    @PostMapping
    public ResponseEntity<ResponseMessage> signUp(@RequestBody UserDTO userDTO) {
        if (Objects.isNull(userDTO.getEmail()))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Invalid Request", null));
        return ResponseEntity.accepted().body(userService.addNewUser(userDTO));
    }

    @PutMapping
    public ResponseEntity<ResponseMessage> updateUser(@RequestHeader("Authorization") String token, @RequestBody UserDTO userDTO) {
        if (Objects.isNull(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        if (!jwtValidationService.validateJwt(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        return ResponseEntity.accepted().body(userService.updateUser(userDTO, token));
    }

    @DeleteMapping("/{userId}")
    public ResponseEntity<ResponseMessage> deleteUser(@RequestHeader("Authorization") String token) {
        if (Objects.isNull(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        if (!jwtValidationService.validateJwt(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        return ResponseEntity.accepted().body(userService.deleteUser(token));
    }

    @PostMapping("/login")
    public ResponseEntity<ResponseMessage> login(@RequestBody LoginRequest loginRequest) {
        if (Objects.isNull(loginRequest.getUsername()) && Objects.isNull(loginRequest.getEmail()))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Invalid Request", null));
        if (Objects.isNull(loginRequest.getPassword()))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Invalid Request", null));
        return ResponseEntity.accepted().body(userService.login(loginRequest));
    }

    @DeleteMapping("/logout")
    public ResponseEntity<ResponseMessage> logout(@RequestHeader("Authorization") String token) {
        if (Objects.isNull(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        if (!jwtValidationService.validateJwt(token))
            return ResponseEntity.badRequest().body(new ResponseMessage("ERROR", "Unauthorized", null));
        return ResponseEntity.accepted().body(userService.logout(token));
    }

}
