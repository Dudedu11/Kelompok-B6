import React, { Component } from 'react';
import {
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
  Image,
} from 'react-native';

export default class Like extends Component {

  state={
    toggle:true
  }

  _onPress(){
    const newState = !this.state.toggle;
    this.setState({toggle:newState})
  }

  render(){
    const {toggle} = this.state;
    const imageValue = toggle?
      <Image
       style={{width:20, height:20}}
       source={require('../assets/icons/love.png')}
      />
       :
      <Image
        style={{width:20, height:20}}
        source={require('../assets/icons/heart.png')}
      />;
    const textValue = toggle?"1,545":"1,546";
    return (
      <View style={styles.container}>
        <View>
          <TouchableOpacity 
          onPress={()=>this._onPress()}
          style={{marginLeft:10, marginTop:210, }}>
            <Text style={{marginTop: 20, fontSize:25, }}>
                {imageValue}
            </Text>
          </TouchableOpacity>
          
          <Text 
            style=
            {{
              color:'black', fontSize:12, position: 'absolute',
              marginLeft:33, marginTop:232.5, 
            }}>
            {textValue}
          </Text>
        </View>
      </View>
    );  
  }
  
};


const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'flex-start',
    },
});