ALTER TABLE `themes` auto_increment = 1;
ALTER TABLE `icons` auto_increment = 1;
ALTER TABLE `reservation_slots` auto_increment = 1;
ALTER TABLE `livestream_tags` auto_increment = 1;
ALTER TABLE `livestream_viewers_history` auto_increment = 1;
ALTER TABLE `livecomment_reports` auto_increment = 1;
ALTER TABLE `ng_words` auto_increment = 1;
ALTER TABLE `reactions` auto_increment = 1;
ALTER TABLE `tags` auto_increment = 1;
ALTER TABLE `livecomments` auto_increment = 1;
ALTER TABLE `livestreams` auto_increment = 1;
ALTER TABLE `users` auto_increment = 1;

USE `isudns`;
ALTER TABLE `records` DROP INDEX domain_id_name_disabled_idx;
ALTER TABLE `records` DROP INDEX name_type_disabled_idx;

ALTER TABLE `records` ADD INDEX domain_id_name_disabled_idx(`domain_id`, `name`, disabled);
ALTER TABLE `records` ADD INDEX name_type_disabled_idx(`name`, `type`, `disabled`);