package com.user.application.password;

import com.user.application.domain.UserReg;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.Objects;

@Component
public class PasswordValidationUtility {

    @Autowired
    PasswordUtility passwordUtility;

    private static final Integer DEFAULT_PASSWORD_VERSION = 2;

    public boolean isPasswordValid(UserReg userReg, String password) {

        return isPasswordValid(userReg, password, DEFAULT_PASSWORD_VERSION);
    }

    public boolean isPasswordValid(UserReg userReg, String password, Integer passwordVersion) {
        Integer version = firstNonNull(passwordVersion, DEFAULT_PASSWORD_VERSION);
        return passwordUtility.checkPassword(password, userReg.getPassword(), version);
    }

    public String hashPassword(String newPassword, Integer passwordVersion) {
        Integer version = firstNonNull(passwordVersion, DEFAULT_PASSWORD_VERSION);
        return passwordUtility.hashPassword(newPassword, version);
    }

    private Integer firstNonNull(Integer passwordVersion, Integer defaultPasswordVersion) {
        return Objects.nonNull(passwordVersion) ? passwordVersion : defaultPasswordVersion;
    }
}

