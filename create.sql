-- 示例文章表
CREATE TABLE `ob_article` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` char(19) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型 1.原创 2.摘抄',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 示例文章表内容表
CREATE TABLE `ob_article_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `article_id` int NOT NULL DEFAULT '0',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` char(19) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `article_id` (`article_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 示例访问表
CREATE TABLE `ob_visits` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `type` tinyint(1) NOT NULL DEFAULT '1',
  `visit_time` char(19) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3066 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 示例数据
INSERT INTO `blog`.`ob_article` (`id`, `title`, `create_time`, `type`) VALUES (125, '话语是一个美丽的陷阱', '2022-01-07 16:45:01', 2);
INSERT INTO `blog`.`ob_article_detail` (`id`, `article_id`, `content`, `create_time`) VALUES (125, 126, '<p>真实是最难的，为了它，一个人也许不得不舍弃许多好东西：名誉，地位，财产，家庭。但真实又是最容易的，在世界上，唯有它，一个人只要愿意，总能得到和保持。 </p><p>人不可能永远真实,也不可能永远虚假。许多真实中有一点虚假，或许多虚假中有一点真实，都是动人的。 </p><p>刻意求真实者太关注自己的形象，已获真实者只是活得自在罢了。 </p><p>活得真诚、独特、潇洒，这样活当然很美。不过，首先要活得自在，才谈得上这些。如果你太关注自己活的样子，总是活给别人看，或者哪怕是活给自己看，那么，你愈是表演得真诚、独特、潇洒，你实际上却活得愈是做作、平庸、拘谨。 </p><p>有的人活得精彩，有的人活得自在，活得潇洒者介乎其间，而非超乎其上。一个人内心生活的隐秘性是在任何情况下都应该受到尊重的，因为隐秘性是内心生活的真实性的保障，从而也是它的存在的保障，内心生活一旦不真实就不复是内心生活了。 </p><p>天赋、才能、眼光、魄力，这一切都还不是伟大，必须加上真实，才成其伟大。真实是一切伟人的共同特征，它源自对人性的真切了解，并由此产生一种面对自己、面对他人的诚实和坦然。 </p><p>精神上的伟人必定是坦诚的，他们足够富有，无需隐瞒自己的欠缺，也足够自尊，不屑于用作秀、演戏、不懂装懂来贬低自己。 </p><p>一个人预先置身于墓中，从死出发来回顾自己的一生，他就会具备一种根本的诚实，因为这时他面对的是自己和上帝。人只有在面对他人时才需要掩饰或撒谎，自欺者所面对的也不是真正的自己，而是自己在他人面前扮演的角色。 </p><p>在不能说真话时，宁愿不说话，也不要说假话。不能说真话而说真话，蠢。不必说假话而说假话，也蠢。 </p><p>撒谎是容易的，带着这谎活下去却是麻烦事，从此你成了它的奴隶，为了圆这谎，你不得不撒更多的也许违背你的心愿且对你有害的谎。</p>', '2022-01-07 16:45:01');