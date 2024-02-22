-- ----------------------------
-- Table structure for draft_tags
-- ----------------------------
DROP TABLE IF EXISTS `draft_tags`;
CREATE TABLE `draft_tags`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `draft_id` int(0) NULL DEFAULT NULL,
  `tag_id` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
