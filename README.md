### QUIZ

#### Features

<ol>
    <li>User</li>
    <li>Category</li>
   
</ol>

#### Requirements

#### Installation

##### User

| Field Name | Data Type   | Default |
| ---------- | ----------- | ------- |
| id         | int         |         |
| first_name | varchar(50) | N/A     |
| last_name  | varchar(50) | N/A     |
| email      | varchar(50) | N/A     |
| username   | varchar(50) | N/A     |
| is_admin   | boolean     | false   |
| is_active  | boolean     | true    |

#### Category table

| Field Name    | Data Type   | Default |
| ----------    | ----------- | ------- |
| id            | int         |         |
| category_name | varchar(50) | N/A     |

#### Question

| Field Name       | Data Type   | Default |
| ----------       | ----------- | ------- |
| id               | int         |         |
| category_id      | int         |         |
| question_title   | text        | N/A     |