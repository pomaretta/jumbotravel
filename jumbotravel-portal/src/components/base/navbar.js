import React, { Fragment } from 'react'
import { Disclosure, Menu, Transition } from '@headlessui/react'
import { BellIcon, MenuIcon, XIcon } from '@heroicons/react/outline'

import { Link } from 'react-router-dom';

import Context from '../context/app';

function classNames(...classes) {
	return classes.filter(Boolean).join(' ')
}

function toSpan({ className, content }) {
	return <span className={className}>{content}</span>
}

class Navbar extends React.Component {

	constructor(props) {
		super(props);

		this.state = {
			currentDate: new Date(),
			interval: null,
		};

	}

	componentDidMount() {
		if (!this.state.interval) {
			this.setState({
				interval: setInterval(() => {
					this.setState({
						currentDate: new Date(),
					});
				}, 60 * 1000),
			});
		}
	}

	componentWillUnmount() {
		clearInterval(this.state.interval);
	}

	render() {
		return (
			<Disclosure as="nav" className="sticky sm:relative bg-white h-14 w-full z-10">
				{({ open }) => (
					<>
						<div className="w-full px-2 sm:px-3 lg:px-5 border-b">
							<div className="relative flex items-center justify-between h-14">
								<div className="absolute inset-y-0 left-0 flex items-center justify-between w-full sm:hidden">
									{/* Mobile menu button*/}
									<Disclosure.Button className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-700 hover:bg-jt-primary focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">
										<span className="sr-only">Open main menu</span>
										{open ? (
											<XIcon className="block h-6 w-6" aria-hidden="true" />
										) : (
											<MenuIcon className="block h-6 w-6" aria-hidden="true" />
										)}
									</Disclosure.Button>
									<lord-icon
										trigger={this.context.newNotifications ? "loop" : ""}
										src="/resources/notification-bell.json"
										style={{
											width: '25px',
											height: '25px',
										}}
									/>
								</div>
								<div className="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
									<div className="hidden sm:block">
										<div className="flex space-x-4">
											<p
												className={classNames(
													'bg-white text-brand-blue font-bold',
													'px-3 py-2 rounded-md text-sm'
												)}
											>
												{
													this.context.agent ?
														`${String(this.context.agent.type).toLowerCase().replace(/\b\w/g, l => l.toUpperCase())} ${this.context.agent.type == "PROVIDER" ? `at ${this.context.agent.getAirportFullName()}` : ''
														}`
														: ''
												}
											</p>
										</div>
									</div>
								</div>
								<div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
									{/* Local Datetime and UTC Datetime */}
									<div className='hidden sm:flex flex-col items-end justify-center'>
										<p className="text-xs text-gray-500">
											<span className='font-bold mr-2'>
												{
													Intl.DateTimeFormat().resolvedOptions().timeZone
												}
											</span>
											{
												this.state.currentDate.toLocaleString().split(' ').slice(0, 5).join(' ').split(':').slice(0, 2).join(':')
											}
										</p>
										<p className="text-xs text-gray-500">
											<span className='font-bold mr-2'>
												UTC
											</span>
											{
												// Get UTC String without seconds
												this.state.currentDate.toUTCString().replace(/GMT.*$/, '').split(' ').slice(0, 5).join(' ').split(':').slice(0, 2).join(':')
											}
										</p>
									</div>
									{/* Profile dropdown */}
									{/* <Menu as="div" className="ml-3 relative">
										<div>
											<Menu.Button className="bg-brand-primary flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-blue-400 focus:ring-white">
												<span className="sr-only">Open user menu</span>
												<lord-icon className="h-8 w-8 rounded-full" trigger="hover" src="/resources/avatar.json" />
											</Menu.Button>
										</div>
										<Transition
											as={Fragment}
											enter="transition ease-out duration-100"
											enterFrom="transform opacity-0 scale-95"
											enterTo="transform opacity-100 scale-100"
											leave="transition ease-in duration-75"
											leaveFrom="transform opacity-100 scale-100"
											leaveTo="transform opacity-0 scale-95"
										>
											<Menu.Items className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
												<Menu.Item>
													{({ active }) => (
														<a
															href="#"
															className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
														>
															Profile
														</a>
													)}
												</Menu.Item>
												<Menu.Item>
													{({ active }) => (
														<a
															href="#"
															className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
														>
															Settings
														</a>
													)}
												</Menu.Item>
											</Menu.Items>
										</Transition>
									</Menu> */}
								</div>
							</div>
						</div>

						<Disclosure.Panel className="sm:hidden | bg-white">
							<div className="px-2 pt-2 pb-3">
								{/* Notifications */}
								<div className={classNames(
									this.context.notificationsIsOpen ? 'mb-2' : 'mb-5',
									'flex items-center justify-between'
								)}
									onClick={this.context.setNotificationsOpen}
								>
									<span className='text-lg flex items-center justify-center'>
										Notifications
										{/* Chevron down */}
										{
											this.context.notificationsIsOpen ? (
												<svg className='h-6 w-6 ml-3 text-gray-700' viewBox='0 0 20 20' fill='currentColor'>
													<path fillRule='evenodd' d='M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z' clipRule='evenodd' />
												</svg>
											) : (
												// Chevron top
												<svg className='h-6 w-6 ml-3 text-gray-700' viewBox='0 0 20 20' fill='currentColor'>
													<path fillRule='evenodd' d='M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z' clipRule='evenodd' />
												</svg>
											)
										}
									</span>
									{/* Notification Bell */}
									{/* <div className="ml-2 rounded-md text-gray-400 hover:text-gray-700 hover:bg-jt-primary focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">
										<lord-icon
											trigger={this.context.newNotifications ? "loop" : ""}
											src="/resources/notification-bell.json"
											style={{
												width: '25px',
												height: '25px',
											}}
										/>
									</div> */}
								</div>

								<div
									className={classNames(
										this.context.notificationsIsOpen ? 'block' : 'hidden',
										'bg-white rounded-md shadow-md w-full h-full p-2 mb-4 mt-2',
									)}
								>
									{/* Notification list */}
									{
										this.context.hasNotifications ? (
											this.context.notifications.getAll().map((notification, index) => {
												return (
													<div
														key={index}
														className="flex flex-row items-center justify-center w-full"
													>
														{
															notification.getNotification()
														}
													</div>
												)
											})
										) : (
											<div className="flex items-center justify-start">
												<lord-icon
													src="/resources/inbox-gray-700.json"
													trigger="loop"
													style={{
														width: '25px',
														height: '25px',
													}}
												/>
												<p className="ml-5 text-sm text-gray-700">
													Nothing to see here.
												</p>
											</div>
										)
									}
								</div>

								{/* Navigation */}
								<Link
									to={'/flights'}
									className={classNames(
										'bg-white text-gray-700',
										'block px-3 py-2 rounded-md text-base font-medium mb-2 mt-3 shadow hover:shadow-md'
									)}
								>
									Flights
								</Link>
								<Link
									to={'/bookings'}
									className={classNames(
										'bg-white text-gray-700',
										'block px-3 py-2 rounded-md text-base font-medium mb-2 shadow hover:shadow-md'
									)}
								>
									Bookings
								</Link>
								<a
									className={classNames(
										'bg-red-500 text-white',
										'block px-3 py-2 rounded-md text-base font-medium mb-2 shadow'
									)}
									onClick={() => {
										this.context.logout()
									}}
								>
									Logout
								</a>
							</div>
						</Disclosure.Panel>
					</>
				)}
			</Disclosure>
		)
	}

}

Navbar.contextType = Context;

export default Navbar;