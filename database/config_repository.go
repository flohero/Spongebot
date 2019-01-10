package database

import "github.com/flohero/Spongebot/database/model"

func (p *Persistence) CreateConfig(conf *model.Config) {
	p.db.Create(&conf)
}

func (p *Persistence) FindConfigById(id int) (conf *model.Config) {
	conf = &model.Config{}
	p.db.Where(&model.Config{Id: id}).First(conf)
	return conf
}

func (p *Persistence) FindFirstActiveConfig() (*model.Config, error) {
	conf := &model.Config{}
	err := p.db.Where(&model.Config{Active: true}).First(conf).Error
	return conf, err
}

func (p *Persistence) FindConfigByToken(token string) (config *model.Config) {
	config = &model.Config{}
	p.db.Where(&model.Config{Token: token}).First(config)
	return config
}

func (p *Persistence) FindAllConfigs() ([]*model.Config, error) {
	rows, err := p.db.Table("configs").Rows()
	if err != nil {
		return nil, err
	}
	configs := make([]*model.Config, 0)
	for rows.Next() {
		conf := &model.Config{}
		if err := rows.Scan(&conf.Id, &conf.Token); err != nil {
			return nil, err
		}
		configs = append(configs, conf)
	}
	return configs, nil
}
