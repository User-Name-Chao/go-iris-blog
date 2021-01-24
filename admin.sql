CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_name` varchar(32) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `status` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `city_name` varchar(12) DEFAULT NULL,
  `city_id` int(11) DEFAULT NULL,
  `admin_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_admin_city_id` (`city_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
 
-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'davie', '2020-05-03 17:05:27', '1', '', '123', '湖北武汉', '1', null);
INSERT INTO `admin` VALUES ('2', 'lili', '2020-05-21 17:05:36', '2', '2', '123', '湖北随州', '2', null);