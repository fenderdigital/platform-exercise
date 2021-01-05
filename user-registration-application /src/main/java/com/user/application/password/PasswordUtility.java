package com.user.application.password;

public interface PasswordUtility {
    String hashPassword(String password, Integer hashVersion);

    boolean checkPassword(String password, String hashedPassword, Integer hashVersion);
}
