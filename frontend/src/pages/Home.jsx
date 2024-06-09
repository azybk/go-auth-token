const Home = ({name}) => {
    return (
        <div>{name ? "hi " + name : "You're not logged in"}</div>
    )
}

export default Home