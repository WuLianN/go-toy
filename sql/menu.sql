
-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '菜单名称',
  `path` varchar(255) NOT NULL COMMENT '菜单路径',
  `component` varchar(255) NULL DEFAULT NULL COMMENT '组件名称',
  `redirect` varchar(255) NULL DEFAULT NULL COMMENT '重定向',
  `parent_id` int(0) UNSIGNED NOT NULL COMMENT '父级id',
  `meta_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT 'meta id',
  `is_use` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '1 使用 2不使用',
  `category` varchar(255) NOT NULL COMMENT '分类名称',
  `sort` int(0) NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for menu_meta
-- ----------------------------
DROP TABLE IF EXISTS `menu_meta`;
CREATE TABLE `menu_meta`  (
  `id` int(0) NOT NULL,
  `title` varchar(255) NULL DEFAULT NULL,
  `hide_menu` int(0) NULL DEFAULT NULL,
  `icon` varchar(255) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL COMMENT '排序',
  `hide_children_in_menu` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;