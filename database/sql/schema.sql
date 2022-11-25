CREATE DATABASE `mr.jackpot`;
USE `mr.jackpot`;



CREATE TABLE user (
    user_id bigint NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (user_id)
);

CREATE TABLE visitor (
    identifier varchar(128) NOT NULL,
    visitor_id bigint       NOT NULL,

    PRIMARY KEY (identifier),
    FOREIGN KEY (visitor_id) REFERENCES user (user_id)
);

CREATE TABLE customer (
    customer_id bigint    NOT NULL,
    status    boolean     NOT NULL DEFAULT TRUE,

    name    varchar(16) NOT NULL,
    address varchar(256),
    phone   varchar(16),

    orders  tinyint NOT NULL DEFAULT 0,
    rating  tinyint NOT NULL DEFAULT 0,
    paid    int     NOT NULL DEFAULT 0,

    PRIMARY KEY (customer_id),
    FOREIGN KEY (customer_id) REFERENCES user (user_id)
);

CREATE TABLE customer_auth (
    id          varchar(32) NOT NULL,
    password    varchar(60) NOT NULL,
    customer_id bigint      NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (customer_id) REFERENCES customer (customer_id)
);



CREATE TABLE coupon_issued (
    coupon_id   bigint      NOT NULL AUTO_INCREMENT,
    code        varchar(64) NOT NULL,
    amount      int         NOT NULL DEFAULT 0,
    title       varchar(256),
    description varchar(256),
    created_at  timestamp NOT NULL,
    expires_at  timestamp NOT NULL,

    PRIMARY KEY (coupon_id),
    UNIQUE KEY (code)
);

CREATE TABLE coupon_owned (
    coupon_id  bigint  NOT NULL,
    owner_id   bigint  NOT NULL,
    valid      boolean NOT NULL DEFAULT FALSE,

    PRIMARY KEY (coupon_id, owner_id),
    FOREIGN KEY (coupon_id) REFERENCES coupon_issued (coupon_id),
    FOREIGN KEY (owner_id)  REFERENCES customer (customer_id)
);



CREATE TABLE entity_type (
    type_id  tinyint     NOT NULL AUTO_INCREMENT,
    typename varchar(64) NOT NULL,

    PRIMARY KEY (type_id)
);

CREATE TABLE pro_order_choice (
    seq_id  tinyint      NOT NULL AUTO_INCREMENT,
    tag     varchar(64)  NOT NULL,
    target  varchar(64)  NOT NULL,
    message varchar(256) NOT NULL,
    type_id tinyint      NOT NULL,

    PRIMARY KEY (seq_id),
    FOREIGN KEY (type_id) REFERENCES  entity_type (type_id)
);

CREATE TABLE pre_order_choice (
    seq_id  tinyint      NOT NULL AUTO_INCREMENT,
    tag     varchar(64)  NOT NULL,
    message varchar(256) NOT NULL,

    PRIMARY KEY (seq_id)
);

CREATE TABLE pro_order_choice_nxt_seq (
    seq_id tinyint NOT NULL,
    nxt_id tinyint NOT NULL,
    
    PRIMARY KEY (seq_id, nxt_id),
    FOREIGN KEY (seq_id) REFERENCES pro_order_choice (seq_id),
    FOREIGN KEY (nxt_id) REFERENCES pre_order_choice (seq_id)
);

CREATE TABLE pre_order_choice_nxt_seq (
    seq_id tinyint NOT NULL,
    nxt_id tinyint NOT NULL,
    
    PRIMARY KEY (seq_id, nxt_id),
    FOREIGN KEY (seq_id) REFERENCES pre_order_choice (seq_id),
    FOREIGN KEY (nxt_id) REFERENCES pro_order_choice (seq_id)
);





CREATE TABLE board_entity (
    entity_id tinyint     NOT NULL AUTO_INCREMENT,
    target_id tinyint     NOT NULL,
    name      varchar(64) NOT NULL,
    tag       varchar(64) NOT NULL,
    price     int         NOT NULL DEFAULT 0,

    PRIMARY KEY (entity_id),
    FOREIGN KEY (target_id) REFERENCES pro_order_choice (seq_id)
);

CREATE TABLE dinner (
    dinner_id tinyint NOT NULL AUTO_INCREMENT,
    entity_id tinyint NOT NULL,

    PRIMARY KEY (dinner_id),
    FOREIGN KEY (entity_id) REFERENCES board_entity (entity_id)
);

CREATE TABLE menu_type (
    id   tinyint     NOT NULL AUTO_INCREMENT,
    name varchar(64) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE menu (
    menu_id      tinyint NOT NULL AUTO_INCREMENT,
    entity_id    tinyint NOT NULL,
    type_id      tinyint NOT NULL,
    option1_name varchar(64) DEFAULT NULL,
    option2_name varchar(64) DEFAULT NULL,

    PRIMARY KEY (menu_id),
    FOREIGN KEY (entity_id) REFERENCES board_entity (entity_id),
    FOREIGN KEY (type_id)   REFERENCES menu_type (id)
);

CREATE TABLE menu_option1 (
    menu_id   tinyint NOT NULL,
    option_id tinyint NOT NULL AUTO_INCREMENT,
    entity_id tinyint NOT NULL,

    PRIMARY KEY (option_id),
    FOREIGN KEY (menu_id)   REFERENCES menu (menu_id),
    FOREIGN KEY (entity_id) REFERENCES board_entity (entity_id)
);

CREATE TABLE menu_option2 (
    menu_id   tinyint NOT NULL,
    option_id tinyint NOT NULL AUTO_INCREMENT,
    entity_id tinyint NOT NULL,

    PRIMARY KEY (option_id),
    FOREIGN KEY (menu_id)   REFERENCES menu (menu_id),
    FOREIGN KEY (entity_id) REFERENCES board_entity (entity_id)
);

CREATE TABLE style (
    style_id    tinyint     NOT NULL AUTO_INCREMENT,
    entity_id   tinyint     NOT NULL,
    description varchar(64) NOT NULL,

    PRIMARY KEY (style_id),
    FOREIGN KEY (entity_id) REFERENCES board_entity (entity_id)
);

CREATE TABLE entity_count (
    count_id  tinyint NOT NULL AUTO_INCREMENT,
    count     tinyint NOT NULL,
    target_id tinyint NOT NULL,

    PRIMARY KEY (count_id),
    FOREIGN KEY (target_id) REFERENCES pro_order_choice (seq_id)
);



CREATE TABLE dinners_menu (
    dinner_id     tinyint NOT NULL,
    menu_id       tinyint NOT NULL,
    default_count tinyint NOT NULL DEFAULT 0,

    PRIMARY KEY (dinner_id, menu_id),
    FOREIGN KEY (dinner_id) REFERENCES dinner (dinner_id),
    FOREIGN KEY (menu_id)   REFERENCES menu (menu_id)
);

CREATE TABLE dinners_style (
    dinner_id tinyint NOT NULL,
    style_id  tinyint NOT NULL,

    PRIMARY KEY (dinner_id, style_id),
    FOREIGN KEY (dinner_id) REFERENCES dinner (dinner_id),
    FOREIGN KEY (style_id)  REFERENCES style (style_id)
);



CREATE TABLE `order` (
    order_id      bigint    NOT NULL AUTO_INCREMENT,
    user_id       bigint    NOT NULL,
    price         int       NOT NULL DEFAULT 0,
    deposit       int       NOT NULL DEFAULT 0,
    discount      int       NOT NULL DEFAULT 0,
    reservated_at timestamp NOT NULL,

    PRIMARY KEY (order_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);

CREATE TABLE state (
    state_id tinyint     NOT NULL AUTO_INCREMENT,
    name     varchar(64) NOT NULL,

    PRIMARY KEY (state_id)
);

CREATE TABLE order_state (
    order_id bigint  NOT NULL,
    state_id tinyint NOT NULL DEFAULT 1,

    PRIMARY KEY (order_id),
    FOREIGN KEY (order_id) REFERENCES `order` (order_id),
    FOREIGN KEY (state_id) REFERENCES state (state_id)
);

CREATE TABLE ordered_dinner (
    id       bigint  NOT NULL AUTO_INCREMENT,
    order_id bigint  NOT NULL,
    style_id tinyint NOT NULL,
    amount   int     NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES `order` (order_id),
    FOREIGN KEY (style_id) REFERENCES style (style_id)
);

CREATE TABLE ordered_menu (
    id          bigint  NOT NULL AUTO_INCREMENT,
    order_id    bigint  NOT NULL,
    dinner_id   bigint  NOT NULL,
    menutype_id tinyint NOT NULL,
    menu_id     tinyint NOT NULL,
    option1_id  tinyint NOT NULL,
    option2_id  tinyint NOT NULL,
    count       tinyint NOT NULL DEFAULT 0,
    price       int,

    PRIMARY KEY (id),
    FOREIGN KEY (order_id)   REFERENCES `order` (order_id),
    FOREIGN KEY (dinner_id)  REFERENCES ordered_dinner (id),
    FOREIGN KEY (menu_id)    REFERENCES menu (menu_id),
    FOREIGN KEY (option1_id) REFERENCES menu_option1 (option_id),
    FOREIGN KEY (option2_id) REFERENCES menu_option2 (option_id)
);

CREATE TABLE delivery_info (
    order_id bigint      NOT NULL,
    name     varchar(16) NOT NULL,
    address  varchar(256) NOT NULL,
    phone    varchar(16) NOT NULL,
    message  varchar(512),

    PRIMARY KEY (order_id),
    FOREIGN KEY (order_id) REFERENCES `order` (order_id)
);



CREATE TABLE role (
    role_id tinyint     NOT NULL AUTO_INCREMENT,
    tag     varchar(64) NOT NULL,
    name    varchar(64) NOT NULL,

    PRIMARY KEY (role_id)
);

CREATE TABLE menu_role (
    menu_id tinyint NOT NULL,
    role_id tinyint NOT NULL,

    PRIMARY KEY (menu_id),
    FOREIGN KEY (menu_id) REFERENCES menu (menu_id),
    FOREIGN KEY (role_id) REFERENCES role (role_id)
);

CREATE TABLE staff (
    staff_id  bigint      NOT NULL AUTO_INCREMENT,
    status    boolean     NOT NULL DEFAULT TRUE,
    role_id   tinyint     NOT NULL,
    name      varchar(16) NOT NULL,
    score     tinyint     NOT NULL DEFAULT 0,

    PRIMARY KEY (staff_id),
    FOREIGN KEY (role_id) REFERENCES role (role_id)
);

CREATE TABLE staff_auth (
    code     varchar(16) NOT NULL,
    staff_id bigint      NOT NULL,

    PRIMARY KEY (code),
    FOREIGN KEY (staff_id) REFERENCES staff (staff_id)
);



CREATE TABLE stock (
    stock_id bigint      NOT NULL AUTO_INCREMENT,
    name     varchar(64) NOT NULL,
    count    int         NOT NULL DEFAULT 0,

    PRIMARY KEY (stock_id),
    UNIQUE KEY (name)
);
