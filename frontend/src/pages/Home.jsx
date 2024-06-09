const Home = (props) => {
    return (
        <div>{props.name ? "hi " + props.name : "You're not logged in"}</div>
    )
}

export default Home