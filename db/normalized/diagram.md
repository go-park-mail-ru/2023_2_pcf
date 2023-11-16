@startuml

Class Balance {
  + ID
  + total_balance DECIMAL DEFAULT 0.0 NOT NULL
  + reserved_balance DECIMAL DEFAULT 0.0 NOT NULL
  + available_balance DECIMAL DEFAULT 0.0 NOT NULL
}

Class User {
  + ID
  + login TEXT NOT NULL UNIQUE
  + password TEXT NOT NULL
  + f_name TEXT DEFAULT NULL
  + l_name TEXT DEFAULT NULL
  + s_name TEXT DEFAULT 'company'
  + avatar TEXT DEFAULT 'default.jpg'
  + balance_id INT
}

Class Ad {
  + ID
  + name TEXT NOT NULL
  + description TEXT DEFAULT NULL
  + website_link TEXT NOT NULL
  + budget DECIMAL DEFAULT 0.0 NOT NULL
  + target_id INT
  + image_link TEXT NOT NULL
  + owner_id INT
}

Class Target {
  + ID
  + name TEXT NOT NULL
  + owner_id INT
  + gender TEXT DEFAULT NULL
  + min_age INT DEFAULT 0
  + max_age INT DEFAULT 127
  + tags TEXT DEFAULT NULL
  + keys TEXT DEFAULT NULL
  + regions TEXT DEFAULT NULL
  + interests TEXT DEFAULT NULL
}

Balance --{ User : "1" balance_id - "1" id
User --{ Ad : "1" owner_id - "1" id
Ad --{ Target : "1" target_id - "1" id
Target --{ User : "1" owner_id - "1" id
@enduml
