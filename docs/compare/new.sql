create table domain
(
    domain_id      varchar(64)              not null
        primary key,
    name           varchar(255)             not null,
    catchall_ind   varchar(1)   default 'N' not null,
    catchall_login varchar(255) default ''  not null,
    constraint ix_domain_name
        unique (name)
)
    collate = utf8mb4_general_ci;

create table mailbox_user
(
    user_id   varchar(64)            not null
        primary key,
    domain_id varchar(64)            not null,
    login     varchar(255)           not null,
    password  varchar(255)           not null,
    admin_ind varchar(1) default 'N' not null,
    constraint ix_mailbox_user
        unique (user_id, login),
    constraint fk_mailbox_user_domain
        foreign key (domain_id) references domain (domain_id)
)
    collate = utf8mb4_general_ci;

create table mailbox
(
    mailbox_id        varchar(64)            not null
        primary key,
    user_id           varchar(64)            not null,
    name              varchar(255)           not null,
    flag_non_existent varchar(1) default 'N' not null,
    flag_no_inferiors varchar(1) default 'N' not null,
    flag_no_select    varchar(1) default 'N' not null,
    flag_marked       varchar(1) default 'N' not null,
    flag_subscribed   varchar(1) default 'N' not null,
    flag_remote       varchar(1) default 'N' not null,
    flag_archive      varchar(1) default 'N' not null,
    flag_drafts       varchar(1) default 'N' not null,
    flag_flagged      varchar(1) default 'N' not null,
    flag_junk         varchar(1) default 'N' not null,
    flag_sent         varchar(1) default 'N' not null,
    flag_trash        varchar(1) default 'N' not null,
    flag_important    varchar(1) default 'N' not null,
    constraint fk_mailbox_mailbox_user
        foreign key (user_id) references mailbox_user (user_id)
)
    collate = utf8mb4_general_ci;

create table message
(
    message_id    varchar(64)                               not null
        primary key,
    mailbox_id    varchar(64)                               not null,
    body          longblob                                  null,
    uid           int           default 0                   not null,
    created_date  datetime      default current_timestamp() not null,
    flag_seen     varchar(1)    default 'N'                 not null,
    flag_answered varchar(1)    default 'N'                 not null,
    flag_flagged  varchar(1)    default 'N'                 not null,
    flag_deleted  varchar(1)    default 'N'                 not null,
    flag_draft    varchar(1)    default 'N'                 not null,
    spam_score    decimal(5, 2) default 0.00                not null,
    constraint fk_message_mailbox
        foreign key (mailbox_id) references mailbox (mailbox_id)
)
    collate = utf8mb4_general_ci;

create table queue
(
    queue_id  varchar(64)  not null
        primary key,
    from_addr varchar(255) not null,
    body      longtext     null
)
    collate = utf8mb4_general_ci;

create table queue_recipient
(
    queue_recipient_id varchar(64)            not null
        primary key,
    queue_id           varchar(64)            not null,
    to_addr            varchar(255)           not null,
    attempts           int        default 0   not null,
    last_attempted_dt  datetime               null,
    success_ind        varchar(1) default 'N' not null,
    constraint fk_qrec_queue
        foreign key (queue_id) references queue (queue_id)
)
    collate = utf8mb4_general_ci;

create table session
(
    session_id  varchar(64)  not null,
    user_id     varchar(64)  not null,
    token       varchar(128) not null,
    expired_ind varchar(1)   not null,
    expiry_date datetime     not null,
    data        mediumtext   not null,
    constraint fk_session_user
        foreign key (user_id) references mailbox_user (user_id)
)
    collate = utf8mb4_general_ci;

