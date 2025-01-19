package config

type Config struct {
  DatabaseName string `env:"DB_NAME"`
  DatabaseHost string `env:"DB_HOST"`
  DatabasePort string `env:"DB_PORT"`
  DatabaseUser string `env:"DB_USER"`
  DatabasePassword string `env:"DB_PASSWORD"`
}

func New() Config {
  var cfg Config
  cfg, err := env.ParseAs[Config]();
  if err != nil {
      return nil, fmt.Errorf("failed to load config: %w", err)
  }
  return &cfg, nil
}
