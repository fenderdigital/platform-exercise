package com.user.application.password.Impl;

import com.user.application.password.PasswordUtility;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.crypto.bcrypt.BCrypt;
import org.springframework.stereotype.Component;

@Component
public class PasswordUtilityImpl implements PasswordUtility {
    private static final Logger LOG = LoggerFactory.getLogger(PasswordUtilityImpl.class);
    private int logRounds;

    @Autowired
    public PasswordUtilityImpl(@Value("${spring.password.utility.bcrypt.log.rounds:5}") int logRounds) {
        this.logRounds = logRounds;
    }

    @Override
    public String hashPassword(String password, Integer hashVersion) {
        String result = null;

        if (password == null) {
            return result;
        }

        switch (hashVersion) {
            case 0:
                result = password;
                break;
            case 1:
                result = BCrypt.hashpw(password.toUpperCase(), BCrypt.gensalt(logRounds));
                break;
            case 2:
                result = BCrypt.hashpw(password, BCrypt.gensalt(logRounds));
                break;
            default:
                LOG.error("Unexpected password version hashVersion=" + hashVersion + ", unable to generate hashed password.");
                break;
        }
        return result;
    }

    @Override
    public boolean checkPassword(String password, String hashedPassword, Integer hashVersion) {
        boolean valid;

        if (password == null || hashedPassword == null) {
            return false;
        }

        switch (hashVersion) {
            case 0:
                valid = hashedPassword.equalsIgnoreCase(password);
                break;
            case 1:
                valid = BCrypt.checkpw(password.toUpperCase(), hashedPassword);
                break;
            case 2:
                valid = BCrypt.checkpw(password, hashedPassword);
                break;
            default:
                valid = false;
                LOG.error("Unexpected password version hashVersion=" + hashVersion + ", unable to validate password.");
                break;
        }

        return valid;
    }

}

