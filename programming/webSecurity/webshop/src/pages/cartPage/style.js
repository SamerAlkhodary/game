import { makeStyles } from "@material-ui/core/styles";
const useStyles = makeStyles((theme) => ({
    list: {
      width: '100%',
      maxWidth: '100ch',
      backgroundColor: theme.palette.background.paper,
      marginTop: theme.spacing(10),
     
    },
    root:{
        width: '100%',
        height:'100%',
        backgroundColor: theme.palette.background.paper,
    },
    large: {
        width: theme.spacing(10),
        height: theme.spacing(10),
        marginRight: theme.spacing(2)
      },
    inline: {
      display: 'inline',
    },
    totalAmount:{
        marginTop: theme.spacing(10),
       

    },
    list2: {
        width: '100%',
        height: 400,
        maxWidth: 300,
        backgroundColor: theme.palette.background.paper,
      },
    
  }));
  export default useStyles;