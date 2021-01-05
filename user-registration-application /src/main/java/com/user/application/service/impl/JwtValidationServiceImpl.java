package com.user.application.service.impl;

import com.user.application.cache.TokenHashMap;
import com.user.application.service.JwtValidationService;
import com.user.application.util.JwtUtility;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.Objects;

@Component
public class JwtValidationServiceImpl implements JwtValidationService {

    @Autowired
    JwtUtility jwtUtility;

    @Autowired
    TokenHashMap tokenHashMap;

    private static final String TOKEN_PREFIX = "Bearer";

    @Override
    public Boolean validateJwt(final String token) {
        Boolean isValid;
        String jwt;
        if (token.startsWith(TOKEN_PREFIX)) {
            jwt = jwtUtility.getJwt(token);
            String userId = jwtUtility.getUserId(jwt);
            isValid = !jwtUtility.isTokenExpired(jwt) && Objects.nonNull(userId) && Objects.nonNull(tokenHashMap.get(Integer.valueOf(userId)));
        } else {
            isValid = false;
        }
        return isValid;
    }
}
