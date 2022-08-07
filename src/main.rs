use actix_web::{
    body::BoxBody,
    dev::ServiceResponse,
    error,
    http::{header::ContentType, StatusCode},
    middleware::{self, ErrorHandlerResponse, ErrorHandlers},
    web, App, Error, HttpResponse, HttpServer, Result, Responder,get, post, web::Data
};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
extern crate csv;
use csv::Reader;
use sailfish::runtime::Buffer;

#[derive(Deserialize)]
struct InputTag {
    tag: String,
}
#[derive(Deserialize, Clone)]
struct MapTag{
    Index: usize,
    Tag: String,
}
#[derive(Deserialize, Clone)]
struct GeoTag{
    Lat: f32,
    Log: f32,
    Time: String,
    Url: String,
}

#[get("/")]
async fn get_tag(
    tagmap: Data<HashMap<String, usize>>,
    vec: Data<Vec<GeoTag>>,//出力データが入った構造体配列
    query: web::Query<InputTag>//tagの呼び出しはquery.tag;で
) -> Result<HttpResponse, Error> {
    let tagg = &query.tag.clone();
    let tag_index:usize = tagmap[tagg];
    let mut ryubuf = ryu::Buffer::new();
    let mut json = Buffer::with_capacity(100000);
    json.push_str(r#"{"tag": ""#);
    json.push_str(tagg);
    json.push_str(r#"","results":["#);
    for i in tag_index..tag_index+99{
        json.push_str(r#"{"lat":"#);
        json.push_str(ryubuf.format(vec[i].Lat));
        json.push_str(r#","lon":"#);
        json.push_str(ryubuf.format(vec[i].Log));
        json.push_str(r#","date":""#);
        json.push_str(&vec[i].Time);
        json.push_str(r#"","url":""#);
        json.push_str(&vec[i].Url);
        json.push_str(r#""},"#);
    }
    let i = tag_index + 99;
    json.push_str(r#"{"lat":"#);
    json.push_str(ryubuf.format(vec[i].Lat));
    json.push_str(r#","lon":"#);
    json.push_str(ryubuf.format(vec[i].Log));
    json.push_str(r#","date":""#);
    json.push_str(&vec[i].Time);
    json.push_str(r#"","url":""#);
    json.push_str(&vec[i].Url);
    json.push_str(r#""}]}"#);
    Ok(HttpResponse::build(StatusCode::OK)
        .content_type("application/json")
        .body(json.into_string()))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let mut tagmap = HashMap::new();
    let mut rdr = Reader::from_path("map.csv")?;
    for result in rdr.deserialize() {
        let record: MapTag = result?;
        let counter = tagmap.entry(record.Tag).or_insert(record.Index);
    }
    let mut vec:Vec<GeoTag> = Vec::new();
    let mut rdr = Reader::from_path("geotag.csv")?;
    for result in rdr.deserialize() {
        let record: GeoTag = result?;
        vec.push(record);
    }
    println!("ready");

    HttpServer::new(move || {    
        App::new()
            .app_data(Data::new(tagmap.clone()))
            .app_data(Data::new(vec.clone()))
            .service(get_tag)
    })
    // ローカルホストのport8080で起動
    .bind("127.0.0.1:8080")?
    .run()
    .await
}