/*
 Navicat Premium Data Transfer

 Source Server         : Go
 Source Server Type    : PostgreSQL
 Source Server Version : 130004
 Source Host           : localhost:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130004
 File Encoding         : 65001

 Date: 30/08/2021 15:21:32
*/


-- ----------------------------
-- Sequence structure for items_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."items_id_seq";
CREATE SEQUENCE "public"."items_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for lists_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."lists_id_seq";
CREATE SEQUENCE "public"."lists_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for items
-- ----------------------------
DROP TABLE IF EXISTS "public"."items";
CREATE TABLE "public"."items" (
  "id" int4 NOT NULL DEFAULT nextval('items_id_seq'::regclass),
  "list_id" int4 NOT NULL,
  "title" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "desc" varchar(255) COLLATE "pg_catalog"."default",
  "priority" int2 DEFAULT 3,
  "cost" numeric(10,2),
  "status" int2 NOT NULL,
  "created_at" int4 NOT NULL,
  "updated_at" int4,
  "deleted_at" int4
)
;

-- ----------------------------
-- Records of items
-- ----------------------------

-- ----------------------------
-- Table structure for lists
-- ----------------------------
DROP TABLE IF EXISTS "public"."lists";
CREATE TABLE "public"."lists" (
  "id" int4 NOT NULL DEFAULT nextval('lists_id_seq'::regclass),
  "title" varchar(150) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "total_cost" numeric(10,2) NOT NULL DEFAULT 0,
  "status" int2 NOT NULL,
  "created_at" int4 NOT NULL,
  "updated_at" int4,
  "deleted_at" int4
)
;

-- ----------------------------
-- Records of lists
-- ----------------------------
INSERT INTO "public"."lists" VALUES (2, 'Baslik', 1, 0.00, 1, 1630099652, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (3, 'Baslik', 1, 0.00, 1, 1630099663, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (4, 'Baslik', 1, 0.00, 1, 1630099733, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (5, 'Baslik 5', 1, 0.00, 1, 1630099744, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (6, 'Baslik 5', 1, 0.00, 1, 1630099750, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (7, 'Baslik 58', 1, 0.00, 1, 1630099816, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (8, 'Baslik 58', 1, 896.77, 1, 1630099828, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (9, 'Baslik 58', 1, 896.77, 1, 1630143964, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (10, 'Baslik 58', 1, 896.77, 1, 1630144468, 1630099367, NULL);
INSERT INTO "public"."lists" VALUES (11, 'newTitle', 1, 896.77, 2, 1630144520, 1630238701, NULL);
INSERT INTO "public"."lists" VALUES (1, 'Baslik', 1, 0.00, 1, 1630099367, 1630099367, 1630239972);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 'omer', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 'omer kaplan', 'kaplanomer@outlook.com');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."items_id_seq"
OWNED BY "public"."items"."id";
SELECT setval('"public"."items_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."lists_id_seq"
OWNED BY "public"."lists"."id";
SELECT setval('"public"."lists_id_seq"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 2, true);

-- ----------------------------
-- Primary Key structure for table items
-- ----------------------------
ALTER TABLE "public"."items" ADD CONSTRAINT "items_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table lists
-- ----------------------------
ALTER TABLE "public"."lists" ADD CONSTRAINT "lists_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
