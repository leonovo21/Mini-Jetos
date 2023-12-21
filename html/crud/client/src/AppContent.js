import  React, {Component} from 'react'


export default class AppContent extends Component{

    state = {posts: []}
    constructor(props){
        super(props)
        this.listRef = React.createRef()
    }
    fetchList = () => {
        fetch('https://jsonplaceholder.typicode.com/posts')
            .then((response) => response.json())
            .then(json => {
                this.setState({posts: json})
            })
    }

    clickdItem = (x) =>{
        console.log("clicked", x)
    }

    render(){
        return(
            <p>
                <br/>
                <hr />
                <button onClick={this.fetchList} class ="btn btn-primary" href='#'>Fetch Data</button>
                <hr />

                <ul>
                    {this.state.posts.map((c) => (
                        <li key={c.id}>
                            <a href='#!' onClick={() => this.clickdItem(c.id)}>
                                {c.title}
                            </a>
                        </li>
                        
                    ))}
                </ul>
            </p>

        )
    }
}
