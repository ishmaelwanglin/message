package feishu

/*
{
    "msg_type": "interactive",
    "card": {
        "config": {
            "wide_screen_mode": true
        },
        "header": {
            "template": "red",
            "title": {
                "tag": "plain_text",
                "content": "👋集群插盘实时异常报警"
            }
        },
        "elements": [
            {
                "tag": "markdown",
                "content": "**磁盘信息**:  1012-1B-B1T-1\n**插盘时间**: 2024-09-18\n**错误信息**:"
            },
            {

                "tag": "hr"
            },
            {
                "tag": "markdown",
                "content": "💡消息来源:DiskMon 😘"
            }
        ]
    }
}
*/

/*
{
  "header": {
    "template": "red",
    "title": {
      "tag": "plain_text",
      "content": "👋集群插盘实时异常报警"
    }
  },
  "elements": [
    {
      "tag": "markdown",
      "content": "**磁盘信息**:  1012-1B-B1T-1\n**插盘时间**: 2024-09-18\n**错误信息**:"
    },
    {
      "tag": "hr"
    },
    {
      "tag": "markdown",
      "content": "💡消息来源:DPL 😘"
    },
    {
      "tag": "column_set",
      "flex_mode": "none",
      "background_style": "grey",
      "columns": [
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        }
      ]
    },
    {
      "tag": "column_set",
      "flex_mode": "none",
      "background_style": "grey",
      "columns": [
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        }
      ]
    },
    {
      "tag": "column_set",
      "flex_mode": "none",
      "background_style": "grey",
      "columns": [
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        },
        {
          "tag": "column",
          "width": "weighted",
          "weight": 1,
          "vertical_align": "top",
          "elements": [
            {
              "tag": "markdown",
              "content": "1012-1B-B1T-1"
            }
          ]
        }
      ]
    }
  ]
}
*/

/*

 "elements": [
    {
      "tag": "div",
      "text": {
        "content": "这是一段普通文本😄",
        "tag": "plain_text"
      }
    },
}
*/

/*
// 7.4 支持
{
    "tag": "table",
    "page_size": 3,
    "row_height": "low",
    "freeze_first_column": true,
    "header_style": {
        "text_align": "left",
        "text_size": "normal",
        "background_style": "none",
        "text_color": "grey",
        "bold": true,
        "lines": 0
    },
    "columns": [
        {
            "name": "first",
            "display_name": "1012-1B-B1T-1",
            "width": "auto",
            "data_type": "text",
            "horizontal_align": "left"
        },
        {
            "name": "second",
            "display_name": "1012-1B-B1T-1",
            "width": "auto",
            "data_type": "text",
            "horizontal_align": "left"
        },
        {
            "name": "third",
            "display_name": "1012-1B-B1T-1",
            "width": "auto",
            "data_type": "text",
            "horizontal_align": "left"
        },
        {
            "name": "fourth",
            "display_name": "1012-1B-B1T-1",
            "width": "auto",
            "data_type": "text",
            "horizontal_align": "left"
        },
    ],
    "rows": [
        {
            "customer_name": "1012-1B-B1T-1",
            "customer_date": 1699341315000,
            "customer_scale": [
                {
                    "text": "S2",
                    "color": "blue"
                }
            ],
            "customer_arr": 168,
            "customer_poc": [
                "ou_14a32f1a02e64944cf19207aa43abcef",
                "ou_e393cf9c22e6e617a4332210d2aabcef"
            ],
            "customer_link": "[飞书科技](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)"
        },
        {
            "customer_name": "飞书科技_01",
            "customer_date": 1606101072000,
            "customer_scale": [
                {
                    "text": "S1",
                    "color": "red"
                }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[飞书科技_01](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"
        },
        {
            "customer_name": "飞书科技_02",
            "customer_date": 1606101072000,
            "customer_scale": [
                {
                    "text": "S3",
                    "color": "orange"
                }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[飞书科技_02](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"

        },
        {
            "customer_name": "飞书科技_03",
            "customer_date": 1606101072000,
            "customer_scale": [
                {
                    "text": "S2",
                    "color": "blue"
                }
            ],
            "customer_arr": 168.23,
            "customer_poc": "ou_14a32f1a02e64944cf19207aa43abcef",
            "customer_link": "[飞书科技_03](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)",
            "company_image": "![image.png](image_key)"
        }
    ]
}


*/
