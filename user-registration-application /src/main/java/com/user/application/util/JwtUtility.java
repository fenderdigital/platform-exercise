package com.user.application.util;

import com.user.application.domain.UserReg;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.MalformedJwtException;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.function.Function;

@Component
public class JwtUtility {

    private static final String SECRET = "secret";
    private static final String TOKEN_PREFIX = "Bearer";


    public String generateToken(UserReg userReg) {
        Map<String, Object> claims = new HashMap<>();
        return createToken(claims, userReg.getUserId().toString());
    }

    //    public String generateToken(UserDetails userDetails){
//        Map<String,Object> claims = new HashMap<>();
//        return createToken(claims, userDetails.getUsername());
//    }

    public <T> T extractClaim(String token, Function<Claims, T> claimsResolver) {
        final Claims claims = extractAllClaims(token);
        return claimsResolver.apply(claims);
    }

    public Date getExpirationDate(String token) {
        try {
            return extractClaim(token, Claims::getExpiration);
        } catch (MalformedJwtException ex) {
            return null;
        }
    }

    public String getUserId(String token) {
        try {
            return extractClaim(token, Claims::getSubject);
        } catch (MalformedJwtException ex) {
            return null;
        }
    }

    public String getJwt(String token) {
        return token.replace(TOKEN_PREFIX, "").trim();
    }

    public Boolean isTokenExpired(String token) {
        try {
            Date date = getExpirationDate(token);
            return Objects.nonNull(date) && date.before(new Date(System.currentTimeMillis()));
        } catch (MalformedJwtException ex) {
            return true;
        }
    }

    private Claims extractAllClaims(String token) {
        return Jwts.parser().setSigningKey(SECRET)
                .parseClaimsJws(token)
                .getBody();
    }


    private String createToken(Map<String, Object> claims, String subject) {
        return Jwts.builder().setClaims(claims).setSubject(subject)
                .setIssuedAt(new Date(System.currentTimeMillis()))
                .setExpiration(new Date(System.currentTimeMillis() + 1000 * 60 * 60))
                .signWith(SignatureAlgorithm.HS256, SECRET).compact();
    }

}

