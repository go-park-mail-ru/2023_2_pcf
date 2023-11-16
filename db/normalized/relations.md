!(//www.plantuml.com/plantuml/png/ZPD1Y-8m5CVl_HJ1YuT5qRqixCckkb3eTXbeOMxbrJn3nZGjIQeZCjzzDCOu7SsCfq9-lxoyV_8wqmQKwMeHpGLeJVw00Bb2SeW8-K6IXVsOne0eAuSMotdoBqwFtxznaUPaCfwGx7zEiY9DhQvGezeYVNm7R84Bg0G-jkKjSjqM6bM_LT4mBac-VCwzJeeiUIgM5hUWzQvHj6zOj2ubr6Y1FzL3yHNKT-1mrTGjoFtmai-0-cHGN4CdpFYbPMVBPIyv9KcMNrF6j9_HdzOFG56l56yDRsIWqHrMcXii1PUREmMgZZ8qtx-M0NNqN9jsXTV0WcMRdKHr4ogtHGB1mWNi0aD9KGL2rboMmF0aNoES2RpUaEdFNouJqu626zo7a4B66ncYN1fKg4sV-n6-_6gZqU4yoB_9O3huE0KZky1fPF79Ywcp_9ru9wRMSFTgBVTIFSrnVsYet0mbxMhn3W00)

**1. Relation "balance"**:

Таблица "balance" с тремя атрибутами: total_balance, reserved_balance и available_balance.
Таблица для хранения балансов юзеров.

Функциональные зависимости:
- Нет дополнительных функциональных зависимостей между атрибутами, так как они все зависят от первичного ключа id.

**Соответствие НФБК**:
- Эта таблица соответствует 1НФ, так как все атрибуты атомарны.
- Соответствует 2НФ, так как у нас нет составных ключей, и первичный ключ id однозначно определяет остальные атрибуты.
- Соответствует 3НФ, так как нет транзитивных зависимостей.

**2. Relation "user"**:

Таблица "user" с несколькими атрибутами, включая login, password и balance_id.
Таблица Юзеров для хранения информации о юзерах. Например: Логин, пароль, баланс, объявления, ФИО и т.д.

Функциональные зависимости:

- `{id} -> login, password, f_name, l_name, s_name, balance_id`
- `{login} -> id`
- `{balance_id} -> total_balance, reserved_balance, available_balance`

**Соответствие НФБК**:
- Эта таблица соответствует 1НФ, так как все атрибуты атомарны.
- Соответствует 2НФ, так как у нас нет составных ключей, и первичный ключ id однозначно определяет остальные атрибуты.
- Соответствует 3НФ, так как нет транзитивных зависимостей.

**3. Relation "ad"**:

Таблица "ad", для хранения информации о Рекламном объявлении. Например: Название, описание, таргетинг, ID владельца.

Функциональные зависимости:
- `{id} -> name, description, website_link, budget, target_id, image_link, owner_id`
- `{owner_id} -> login, password, f_name, l_name, s_name, balance_id, total_balance, reserved_balance, available_balance`
- `{target_id} -> name, owner_id, gender, min_age, max_age, regions, interests, keys, tags`

**Соответствие НФБК**:
- Эта таблица соответствует 1НФ, так как все атрибуты атомарны.
- Соответствует 2НФ, так как у нас нет составных ключей, и первичный ключ id однозначно определяет остальные атрибуты.
- Соответствует 3НФ, так как нет транзитивных зависимостей.

**4. Relation "target"**:

Таблица "target" для хранения информации о таргетингах: Целевой возраст, регионы, пол

Функциональные зависимости:
- `{id} -> name, owner_id, gender, min_age, max_age, regions, interests, keys, tags`

**Соответствие НФБК**:
- Эта таблица соответствует 1НФ, так как все

 атрибуты атомарны.
- Соответствует 2НФ, так как у нас нет составных ключей, и первичный ключ id однозначно определяет остальные атрибуты.
- Соответствует 3НФ, так как нет транзитивных зависимостей.
