use actix_web::{
    body::BoxBody,
    dev::ServiceResponse,
    error,
    http::{header::ContentType, StatusCode},
    middleware::{self, ErrorHandlerResponse, ErrorHandlers},
    web, App, Error, HttpResponse, HttpServer, Result, Responder,get, post
};
use serde::{Deserialize, Serialize};


#[derive(Deserialize)]
struct InputTag {
    tag: String,
}

#[derive(Deserialize)]
struct UserResponse{
    lat: String,
    lon: String,
    date: String,
    url: String,
}

#[get("/")]
async fn get_tag(
    query: web::Query<InputTag>//tagの呼び出しはquery.tag;で
) -> Result<HttpResponse, Error> {

    Ok(HttpResponse::Ok().body(format!("tag is {}", tag)))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {    
        App::new()
            .service(get_tag)
    })
    // ローカルホストのport8080で起動
    .bind("127.0.0.1:8080")?
    .run()
    .await
}