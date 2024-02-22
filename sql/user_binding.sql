-- ----------------------------
-- Table structure for user_binding
-- ----------------------------
DROP TABLE IF EXISTS `user_binding`;
CREATE TABLE `user_binding`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id_1` int(0) NOT NULL,
  `user_id_2` int(0) NOT NULL,
  `created_at` datetime(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_binding_user1`(`user_id_1`, `user_id_2`) USING BTREE,
  UNIQUE INDEX `unique_binding_user2`(`user_id_2`, `user_id_1`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
