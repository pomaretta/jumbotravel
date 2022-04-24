import { Component } from 'react';
import Context from '../context/app';
import ClassName from '../utils/classname';

function ActionDelete(props) {
    return (
        <div className='w-full h-full | flex justify-center items-center'>
            <div className="relative p-4 w-full max-w-md h-full md:h-auto">
                <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
                    <div className="flex justify-end p-2">
                        <button type="button" className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white" data-modal-toggle="popup-modal">
                            <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                        </button>
                    </div>
                    <div className="p-6 pt-0 text-center">
                        <svg className="mx-auto mb-4 w-14 h-14 text-gray-400 dark:text-gray-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                        <h3 className="mb-5 text-2xl sm:text-lg font-normal text-gray-500 dark:text-gray-400">
                            {
                                props.data && props.data.title ? props.data.title : null
                            }
                        </h3>
                        <button 
                            type="button" className="sm:w-auto w-full text-center mb-5 sm:mb-0 text-white bg-red-600 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg | text-2xl sm:text-sm | justify-center sm:justify-start sm:inline-flex | items-center px-5 py-2.5 mr-2"
                            onClick={() => {
                                props.changeState({
                                    show: false,
                                    pressed: {
                                        status: "completed",
                                        proceed: true,
                                    }
                                })
                            }}
                        >
                            Yes, I'm sure
                        </button>
                        <button 
                            type="button" className="sm:w-auto w-full text-center text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-gray-200 rounded-lg | text-2xl sm:text-sm | justify-center sm:justify-start sm:inline-flex | border border-gray-200 font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10"
                            onClick={() => {
                                props.changeState({
                                    show: false,
                                    pressed: {
                                        status: "rejected",
                                        proceed: false,
                                    }
                                })
                            }}
                        >No, cancel</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

function getProductImage(productCode) {
    switch (productCode) {
        case 1:
            return "/images/0001-coca-cola.png";
        case 2:
            return "/images/0002-agua.png";
        case 3:
            return "/images/0003-coca-cola-light.webp";
        case 4:
            return "/images/0004-wine.jpeg";
        case 5:
            return "/images/0005-lays.png";
        case 6:
            return "/images/0006-aceitunas.jpeg";
        case 7:
            return "/images/0007-fanta-naranja.png";
        case 8:
            return "/images/0008-fanta-limon.jpeg";
        case 9:
            return "/images/0009-frutos-secos.jpeg";
        case 10:
            return "/images/0010-galletas.jpeg";
        case 11:
            return "/images/0011-cerveza.jpeg";
        default:
            return "/images/0000-default.svg";
    }
}

function getProductName(products, productCode) {
    const product = products.find(product => product.product_code === productCode);
    if (product) {
        return product.name;
    }
    return "";
}

function ActionPlaceOrder(props) {
    return (
        <div className='w-full h-full | flex justify-center items-center'>
            <div className="relative p-4 w-full h-full | sm:flex sm:justify-center sm:items-center">
                <div className="relative bg-white rounded-lg shadow dark:bg-gray-700 | h-full sm:w-1/2 sm:h-auto | flex flex-col justify-between sm:justify-center sm:items-center">
                    <div className='flex w-full flex-col'>
                        <div className="flex justify-between items-start p-5 rounded-t border-b dark:border-gray-600">
                            <h3 className="text-xl font-semibold text-gray-900 lg:text-2xl dark:text-white">
                                Order Summary
                            </h3>
                            <button
                                type="button" className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="defaultModal"
                                onClick={() => {
                                    props.changeState({
                                        show: false,
                                        pressed: {
                                            status: "completed",
                                            proceed: false,
                                        }
                                    })
                                }}
                            >
                                <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                            </button>
                        </div>
                        <div className="p-4 space-y-4 sm:max-h-60 | sm:overflow-scroll | sm:no-scrollbar">
                            {/* <div className='flex justify-between items-center | w-full | bg-gray-50 | p-3 | rounded-md | shadow'>
                            <p>
                                Items: {props.data.items.length}
                            </p>
                        </div> */}
                            {
                                props.data && props.data.items ?
                                    props.data.items.map((item, index) => {
                                        return (
                                            <div className='w-full bg-gray-50 shadow | flex flex-col sm:flex-row justify-between items-center | sm:h-12 | rounded-md | p-2' key={index}>
                                                <div
                                                    className="w-full | flex | justify-between items-start sm:items-center | sm:w-1/3"
                                                >
                                                    <div className="flex | w-full sm:w-auto | justify-start sm:justify-center items-center | space-x-4">
                                                        <img
                                                            className="rounded-md"
                                                            style={{
                                                                width: "35px",
                                                                height: "35px",
                                                            }}
                                                            src={getProductImage(item.product_code)}
                                                        />
                                                        <p className="text-xl sm:text-xs | font-bold | text-brand-blue">
                                                            {
                                                                props.data.products ? getProductName(props.data.products, item.product_code) : null
                                                            }
                                                        </p>
                                                    </div>
                                                    <div className='flex sm:hidden | justify-end items-center | space-x-4'>
                                                        <p className='sm:text-sm text-md | text-gray-700'>
                                                            Items
                                                        </p>
                                                        <p
                                                            className='sm:text-sm text-md | font-semibold | text-gray-600 | px-2 p-1 | bg-blue-200 | rounded-lg | text-center'
                                                        >
                                                            {
                                                                item.quantity
                                                            }
                                                        </p>
                                                    </div>
                                                </div>
                                                <div className="flex | w-full | justify-center sm:justify-end items-center | space-x-3">
                                                    <div className="hidden sm:flex justify-center sm:justify-end items-center space-x-3 pr-2">
                                                        <p className='text-sm | text-gray-700'>
                                                            Items
                                                        </p>
                                                        <p
                                                            className='text-sm | font-semibold | text-gray-600 | px-2 p-1 | bg-blue-200 | rounded-lg | text-center'
                                                        >
                                                            {
                                                                item.quantity
                                                            }
                                                        </p>
                                                    </div>
                                                </div>
                                            </div>
                                        )
                                    }) : null
                            }
                        </div>
                    </div>
                    <div className="flex flex-col sm:flex-row | justify-left w-full items-center p-6 space-y-4 sm:space-y-0 sm:space-x-2 rounded-b border-t border-gray-200 dark:border-gray-600">
                        <button
                            type="button" className="text-white | w-full sm:w-auto | bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg | text-xl sm:text-sm | px-5 py-2.5 text-center"
                            onClick={() => {
                                props.changeState({
                                    show: false,
                                    pressed: {
                                        status: "completed",
                                        proceed: true,
                                    }
                                })
                            }}
                        >Place order</button>
                        <button
                            type="button" className="text-gray-500 | w-full sm:w-auto | bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-blue-300 rounded-lg border border-gray-200 | text-xl sm:text-sm | font-medium px-5 py-2.5 text-center"
                            onClick={() => {
                                props.changeState({
                                    show: false,
                                    pressed: {
                                        status: "completed",
                                        proceed: false,
                                    }
                                })
                            }}
                        >Cancel</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

class Modal extends Component {

    getModal() {
        if (!this.context.show || !(this.context.data && this.context.data.type)) {
            return null;
        }
        switch (String(this.context.data.type).toLowerCase()) {
            case "actionplaceorder":
                return <ActionPlaceOrder data={this.context.data} changeState={this.context.changeState} />
            case "actiondelete":
                return <ActionDelete data={this.context.data} changeState={this.context.changeState} />;
            default:
                return null;
        }
    }

    render() {

        return (
            <div className={ClassName(
                this.context.show ? 'fixed' : 'hidden',
                "z-40 | top-0 left-0 | w-screen h-screen"
            )}>
                <div
                    className='absolute | top-0 left-0 | bg-gray-700 opacity-40 | w-full h-full'
                    onClick={() => {
                        this.context.changeState({
                            show: false,
                            pressed: {
                                status: "rejected",
                            }
                        })
                    }}
                ></div>
                {
                    this.context.show ?
                        this.getModal()
                        : null
                }
            </div>
        )
    }

}
Modal.contextType = Context;

export default Modal;