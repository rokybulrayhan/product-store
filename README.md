in this project i am using 
language : golang
framework : echo
orm : bun
database : postgresql

run project : go run .


i was attached database table query in db/migrations folder path you can use go-migrate for execute this query otherwise directly execute .



api documention is 
1.all crud 
brand:


post 
localhost:8082/product/store/brand
req body:
{
    "Name" : "pran"
}
res body:
{
    "success": true,
    "data": {
        "id": 1,
        "name": "pran",
        "status_id": true,
        "created_at": "2023-11-18T09:42:27.666282Z",
        "updated_at": null,
        "created_by": "",
        "updated_by": ""
    }
}

put 
localhost:8082/product/store/brand/1
req body
{
    "Name" : "pran",
    "status_id" : true
}
res body
{
    "success": true,
    "data": {
        "id": 1,
        "name": "pran",
        "status_id": true,
        "created_at": "2023-11-18T09:42:27.666282Z",
        "updated_at": null,
        "created_by": "",
        "updated_by": ""
    }
}

Get
localhost:8082/product/store/brand?status_id=1(if you use status_id = 2 you got all inactive data . if you didn't us status_id than you get active inactive data)
for active data
{
    "success": true,
    "data": {
        "total": 1,
        "pages": 1,
        "Brands": [
            {
                "id": 1,
                "name": "pran",
                "status_id": true,
                "created_at": "2023-11-18T09:42:27.666282Z",
                "updated_at": null,
                "created_by": "",
                "updated_by": ""
            }
        ]
    }
}

get single body
Get
localhost:8082/product/store/brand/1
{
    "success": true,
    "data": {
        "id": 1,
        "name": "pran",
        "status_id": true,
        "created_at": "2023-11-18T09:42:27.666282Z",
        "updated_at": null,
        "created_by": "",
        "updated_by": ""
    }
}

DElETE
localhost:8082/product/store/brand/1
res body :
{
    "success": true,
    "message": "deleted successfully."
}


product :

using transaction for this api

Post
req body :
{
    "name": "mango",
    "description": "description",
    "specifications": "specifications",
    "brand_id": 1,
    "category_id": 1,
    "supplier_id": 1,
    "unit_price": 67,
    "discount_price": 23,
    "tags": "p1",
    "product_stock": {
        "stock_quantity": 3
    }
}
res body :

    {
    "success": true,
    "data": {
        "id": 2,
        "name": "mango",
        "description": "description",
        "specifications": "specifications",
        "brand_id": 1,
        "category_id": 1,
        "supplier_id": 1,
        "unit_price": 67,
        "discount_price": 23,
        "tags": "p1",
        "status_id": true,
        "created_at": "2023-11-18T09:53:37.737682Z",
        "updated_at": null,
        "created_by": "",
        "updated_by": ""
    }
}
unit price wise sorted
error handled product quantiy less than 




note : simillarly all are same


2.product filter

GET
localhost:8082/product/store/product?max_price=3&min_price=69
for find your unit price wise list

localhost:8082/product/store/product?brand_id=2,1 
search multiple brand

localhost:8082/product/store/product?category_id=1&supplier_id=1&name=mango
search category, supplier, name all simillar

also add pagination for all

unit price wise sorted


3.category based sequence 

localhost:8082/product/store/category/sequence
my initial response body 
it works secuencely wise and maintain parent child rules

{
    "success": true,
    "data": [
        {
            "id": 1,
            "name": "mobile",
            "parent_id": 0,
            "status_id": true,
            "sequence": 1,
            "created_at": "2023-11-18T09:33:30.101022Z",
            "children": [
                {
                    "id": 2,
                    "name": "ios",
                    "parent_id": 1,
                    "status_id": true,
                    "sequence": 1,
                    "created_at": "2023-11-18T09:34:02.19541Z",
                    "children": null,
                    "updated_at": null,
                    "created_by": "",
                    "updated_by": ""
                },
                {
                    "id": 3,
                    "name": "android",
                    "parent_id": 1,
                    "status_id": true,
                    "sequence": 1,
                    "created_at": "2023-11-18T09:35:16.449921Z",
                    "children": null,
                    "updated_at": null,
                    "created_by": "",
                    "updated_by": ""
                }
            ],
            "updated_at": null,
            "created_by": "",
            "updated_by": ""
        },
        {
            "id": 4,
            "name": "watch",
            "parent_id": 0,
            "status_id": true,
            "sequence": 2,
            "created_at": "2023-11-18T09:35:43.22759Z",
            "children": [
                {
                    "id": 5,
                    "name": "smart watch",
                    "parent_id": 4,
                    "status_id": true,
                    "sequence": 2,
                    "created_at": "2023-11-18T09:36:00.395829Z",
                    "children": null,
                    "updated_at": null,
                    "created_by": "",
                    "updated_by": ""
                }
            ],
            "updated_at": null,
            "created_by": "",
            "updated_by": ""
        }
    ]
}


