/*



*/


use warp::Filter;

#[tokio::main]


async fn main() {

    let register = warp::path("register")
        .map(|| "Hello from register");

    let login = warp::path("login")
        .map(|| "Hello from login");

    let logout = warp::path("logout")
        .map(|| "Hello from logout");

    let routes = register
        .or(login)
        .or(logout);

    warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
}
