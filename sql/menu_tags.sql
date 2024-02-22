-- ----------------------------
-- Table structure for menu_tags
-- ----------------------------
DROP TABLE IF EXISTS `menu_tags`;
CREATE TABLE `menu_tags`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int(0) NOT NULL,
  `menu_id` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

