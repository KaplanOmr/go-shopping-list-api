/*
 Navicat Premium Data Transfer

 Source Server         : Local - pg
 Source Server Type    : PostgreSQL
 Source Server Version : 130004
 Source Host           : localhost:5432
 Source Catalog        : sladb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130004
 File Encoding         : 65001

 Date: 05/10/2021 22:05:07
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
-- Sequence structure for items_id_seq1
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."items_id_seq1";
CREATE SEQUENCE "public"."items_id_seq1" 
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
-- Sequence structure for lists_id_seq1
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."lists_id_seq1";
CREATE SEQUENCE "public"."lists_id_seq1" 
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
-- Sequence structure for users_id_seq1
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq1";
CREATE SEQUENCE "public"."users_id_seq1" 
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
  "id" int4 NOT NULL DEFAULT nextval('items_id_seq1'::regclass),
  "list_id" int4 NOT NULL,
  "title" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "priority" int2 DEFAULT 3,
  "cost" numeric(10,2),
  "status" int2 NOT NULL,
  "created_at" int4 NOT NULL,
  "updated_at" int4,
  "deleted_at" int4
)
;

-- ----------------------------
-- Table structure for lists
-- ----------------------------
DROP TABLE IF EXISTS "public"."lists";
CREATE TABLE "public"."lists" (
  "id" int4 NOT NULL DEFAULT nextval('lists_id_seq1'::regclass),
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
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq1'::regclass),
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."items_id_seq"
OWNED BY "public"."items"."id";
SELECT setval('"public"."items_id_seq"', 3, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."items_id_seq1"
OWNED BY "public"."items"."id";
SELECT setval('"public"."items_id_seq1"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."lists_id_seq"
OWNED BY "public"."lists"."id";
SELECT setval('"public"."lists_id_seq"', 13, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."lists_id_seq1"
OWNED BY "public"."lists"."id";
SELECT setval('"public"."lists_id_seq1"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq1"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq1"', 2, false);

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
