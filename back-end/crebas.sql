/*==============================================================*/
/* DBMS name:      PostgreSQL 8                                 */
/* Created on:     12/13/2021 3:18:32 PM                        */
/*==============================================================*/


drop index MENGOMENTARI_FK;

drop index MEMPOSTING_FK;

drop index COMMENT_PK;

drop table COMMENT;

drop index FEED_PK;

drop table FEED;

drop index MENYUKAI2_FK;

drop index MENYUKAI_FK;

drop index MENYUKAI_PK;

drop table MENYUKAI;

drop index USERS_PK;

drop table USERS;

/*==============================================================*/
/* Table: COMMENT                                               */
/*==============================================================*/
create table COMMENT (
   COMMENT_ID           INT4                 not null,
   FEED_ID              INT4                 null,
   USER_ID              INT4                 null,
   COMMENT_TEXT         VARCHAR(1024)        null,
   POST_COMMENT         DATE                 null,
   constraint PK_COMMENT primary key (COMMENT_ID)
);

/*==============================================================*/
/* Index: COMMENT_PK                                            */
/*==============================================================*/
create unique index COMMENT_PK on COMMENT (
COMMENT_ID
);

/*==============================================================*/
/* Index: MEMPOSTING_FK                                         */
/*==============================================================*/
create  index MEMPOSTING_FK on COMMENT (
USER_ID
);

/*==============================================================*/
/* Index: MENGOMENTARI_FK                                       */
/*==============================================================*/
create  index MENGOMENTARI_FK on COMMENT (
FEED_ID
);

/*==============================================================*/
/* Table: FEED                                                  */
/*==============================================================*/
create table FEED (
   FEED_ID              INT4                 not null,
   IMAGE_FEED           VARCHAR(255)         null,
   CAPTION_FEED         VARCHAR(500)         null,
   POST_DATE            DATE                 null,
   constraint PK_FEED primary key (FEED_ID)
);

/*==============================================================*/
/* Index: FEED_PK                                               */
/*==============================================================*/
create unique index FEED_PK on FEED (
FEED_ID
);

/*==============================================================*/
/* Table: MENYUKAI                                              */
/*==============================================================*/
create table MENYUKAI (
   USER_ID              INT4                 not null,
   FEED_ID              INT4                 not null,
   STATUS_LIKE          BOOL                 null,
   TGL_LIKE             TIMESTAMP            null,
   constraint PK_MENYUKAI primary key (USER_ID, FEED_ID)
);

/*==============================================================*/
/* Index: MENYUKAI_PK                                           */
/*==============================================================*/
create unique index MENYUKAI_PK on MENYUKAI (
USER_ID,
FEED_ID
);

/*==============================================================*/
/* Index: MENYUKAI_FK                                           */
/*==============================================================*/
create  index MENYUKAI_FK on MENYUKAI (
USER_ID
);

/*==============================================================*/
/* Index: MENYUKAI2_FK                                          */
/*==============================================================*/
create  index MENYUKAI2_FK on MENYUKAI (
FEED_ID
);

/*==============================================================*/
/* Table: USERS                                                 */
/*==============================================================*/
create table USERS (
   USER_ID              INT4                 not null,
   USERNAME             VARCHAR(20)          null,
   EMAIL                VARCHAR(50)          null,
   FIRSTNAME            VARCHAR(20)          null,
   LASTNAME             VARCHAR(20)          null,
   PHONENUMBER          VARCHAR(13)          null,
   PASSWORD             VARCHAR(60)          null,
   CREATED_DATE         DATE                 null,
   UPDATED_DATE         DATE                 null,
   EMAIL_VERIFIED       BOOL                 null,
   IMAGE_FILE           CHAR(254)            null,
   IDENTITY_TYPE        INT8                 null,
   IDENTITY_NO          VARCHAR(3)           null,
   EMERGENCY_CALL       VARCHAR(16)          null,
   ADDRESS_KTP          VARCHAR(13)          null,
   DOMISILI             VARCHAR(255)         null,
   USER_PKEY            VARCHAR(255)         null,
   constraint PK_USERS primary key (USER_ID)
);

/*==============================================================*/
/* Index: USERS_PK                                              */
/*==============================================================*/
create unique index USERS_PK on USERS (
USER_ID
);

alter table COMMENT
   add constraint FK_COMMENT_MEMPOSTIN_USERS foreign key (USER_ID)
      references USERS (USER_ID)
      on delete restrict on update restrict;

alter table COMMENT
   add constraint FK_COMMENT_MENGOMENT_FEED foreign key (FEED_ID)
      references FEED (FEED_ID)
      on delete restrict on update restrict;

alter table MENYUKAI
   add constraint FK_MENYUKAI_MENYUKAI_USERS foreign key (USER_ID)
      references USERS (USER_ID)
      on delete restrict on update restrict;

alter table MENYUKAI
   add constraint FK_MENYUKAI_MENYUKAI2_FEED foreign key (FEED_ID)
      references FEED (FEED_ID)
      on delete restrict on update restrict;

