import React from 'react'
import './App.css'
import { List, Card } from 'antd'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { loadProducts as loadProductsAction } from 'actions/product_action'
const { Meta } = Card

class App extends React.Component {
  constructor(props) {
    super(props)
    const { loadProducts } = this.props
    loadProducts({ page: 1, limit: 100 })
  }

  renderCover(title, imageUrl) {
    return <img alt={title} src={imageUrl} />
  }

  render() {
    const { loading, products } = this.props
    return (
      <List
        grid={{ gutter: 16, column: 4 }}
        dataSource={products}
        renderItem={item => (
          <List.Item>
            <Card
              hoverable
              style={{ width: 240 }}
              cover={this.renderCover(item.name, item.image)}
            >
              <Meta title={item.name} description={item.name} />
            </Card>
          </List.Item>
        )}
      />
    );
  }
}

App.propTypes = {
  loading: PropTypes.bool.isRequired,
  products: PropTypes.arrayOf(
    PropTypes.object
  )
}

const mapStateToProps = ({
  product: { loading, products }
}) => ({
  loading,
  products
})

export default connect(mapStateToProps, {
  loadProducts: loadProductsAction,
})(App)
