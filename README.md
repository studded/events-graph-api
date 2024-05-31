# TODO:
- Use Dataloader to fix N+1 queries
- Add more pagination - currently on Events and Users list
- Validate date/times (end should be after start, etc.)
- Last admin role for an event should not be deleted

# Example:
```graphql
query showEvent($id: ID!) {
  event(id: $id) {
    id
    name
    location
    description
    startDate
    endDate
    activities {
      name
      description
      startTime
      endTime
    }
    roles {
      user {
        name
      }
      type
    }
    expenses {
      name
      cost
      category
    }
    expenseTotal
    expenseCategories {
      category
      total
    }
  }
}
```
Output:
```json
{
  "data": {
    "event": {
      "id": "1",
      "name": "Birthday",
      "location": "Dublin",
      "description": "Happy Birthday!",
      "startDate": "2024-06-26T00:00:00Z",
      "endDate": "2024-06-26T00:00:00Z",
      "activities": [
        {
          "name": "Cake",
          "description": "Let them eat cake!",
          "startTime": "2024-06-26T00:00:00-07:00",
          "endTime": "2024-06-26T00:00:00-07:00"
        },
        {
          "name": "Presents",
          "description": "Open them up!",
          "startTime": "2024-06-26T00:00:00-07:00",
          "endTime": "2024-06-26T00:00:00-07:00"
        }
      ],
      "roles": [
        {
          "user": {
            "name": "Kevin"
          },
          "type": "admin"
        },
        {
          "user": {
            "name": "Brandon"
          },
          "type": "contributor"
        },
        {
          "user": {
            "name": "Alana"
          },
          "type": "attendee"
        }
      ],
      "expenses": [
        {
          "name": "Snacks",
          "cost": 9.99,
          "category": "food"
        },
        {
          "name": "Drinks",
          "cost": 5.99,
          "category": "food"
        },
        {
          "name": "Balloons",
          "cost": 3.5,
          "category": "decorations"
        }
      ],
      "expenseTotal": 19.48,
      "expenseCategories": [
        {
          "category": "decorations",
          "total": 3.5
        },
        {
          "category": "food",
          "total": 15.98
        }
      ]
    }
  }
}
```