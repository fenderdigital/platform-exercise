package com.user.application.repository;

import com.user.application.domain.UserReg;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<UserReg, Integer> {
    Optional<UserReg> findByEmail(@Param("email") final String email);

    Optional<UserReg> findByUsername(@Param("username") final String username);


}
