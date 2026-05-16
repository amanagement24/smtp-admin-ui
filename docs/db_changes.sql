CREATE TABLE smtp.user_session
(
    session_id   VARCHAR(64) COLLATE utf8mb4_general_ci   NOT NULL,
    user_id      VARCHAR(64) COLLATE utf8mb4_general_ci   NOT NULL,
    token        VARCHAR(255) COLLATE utf8mb4_general_ci  NOT NULL,
    session_data MEDIUMTEXT COLLATE utf8mb4_general_ci    NULL,
    created_date DATETIME                                    NOT NULL,
    expired_ind  VARCHAR(1) COLLATE utf8mb4_general_ci    NOT NULL DEFAULT 'N',

    -- Primary Key Constraint
    CONSTRAINT pk_user_session
        PRIMARY KEY (session_id),

    -- Foreign Key Constraint (Matches mailbox_user exactly)
    CONSTRAINT fk_session_user
        FOREIGN KEY (user_id) REFERENCES smtp.mailbox_user (user_id)
)
    COLLATE = utf8mb4_general_ci;

-- Index for optimized foreign key lookups and queries
CREATE INDEX ix_session_user
    ON smtp.user_session (user_id);

alter table mailbox_user add admin_ind varchar(1) not null default 'N';
