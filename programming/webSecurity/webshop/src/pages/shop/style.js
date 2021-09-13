import { makeStyles } from "@material-ui/core/styles";
const useStyles = makeStyles((theme)=>({
    container:{
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(8,0,6)
    },
    appBar:{
        background:theme.palette.background.paper,
    },
    icon:{
        marginRight: "20px",

    },
    title:{
        
        width: "80px"
    },
    space:{
        flexGrow: 1,
    },

    loginButton:{
        color:'black',
       
    },
    root: {
        
        background: theme.palette.background.paper,
      },
      tile:{
          maxWidth: "200px",
          height: "200px",
          margin: theme.spacing(1,3)
      },
      horizentalList:{
        marginTop: theme.spacing(8),


      },
      list: {
        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'space-around',
        overflow: 'hidden',
        backgroundColor: theme.palette.background.paper,
        
        marginTop: theme.spacing(5)
      },
      tileTitle: {  
        color: 'white',
      },
      
      gridList: {
        flexWrap: 'nowrap',
        // Promote the list into his own layer on Chrome. This cost memory but helps keeping high FPS.
        transform: 'translateZ(0)',
      },
      paper: {
        height: 140,
        width: 100,
      },
    
 
}));
export default useStyles;