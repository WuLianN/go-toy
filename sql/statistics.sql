-- ----------------------------
-- Table structure for statistics_visit
-- ----------------------------
DROP TABLE IF EXISTS `statistics_visit`;
CREATE TABLE `statistics_visit`  (
  `visit_time` datetime(0) NULL DEFAULT NULL,
  `ip` varchar(255) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;