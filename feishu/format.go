package feishu

var Table = `{
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
                "tag": "table",
                "page_size": 4,
                "row_height": "low",
                "freeze_first_column": true,
                "header_style": {
                    "text_align": "left",
                    "text_size": "normal",
                    "background_style": "grey",
                    "text_color": "default",
                    "bold": true,
                    "lines": 1
                },
                "columns": [
                    {
                        "name": "first",
                        "display_name": "前面板示意图",
                        "width": "auto", 
                        "data_type": "text",
                        "horizontal_align": "left"
                    },
                    {
                        "name": "second",
                        "display_name": "****",
                        "width": "auto",
                        "data_type": "text",
                        "horizontal_align": "left"
                    },
                    {
                        "name": "third",
                        "display_name": "****",
                        "width": "auto",
                        "data_type": "text",
                        "horizontal_align": "left"
                    },
                    {
                        "name": "fourth",
                        "display_name": "****",
                        "width": "auto",
                        "data_type": "text",
                        "horizontal_align": "left"
                    }
                ],
                "rows": [
                    {"first": "alian", "second": "","third":"10101", "fourth":"32"},
                    {"first": "alian", "second": "","third":"10101", "fourth":"32"},
                    {"first": "alian", "second": "","third":"10101", "fourth":"32"}
                ]
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
}`
