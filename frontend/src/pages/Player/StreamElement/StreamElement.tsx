/// <reference types="vite-plugin-svgr/client" />

import { EventType } from '@highlight-run/rrweb'
import WebVitalSimpleRenderer from '@pages/Player/StreamElement/Renderers/WebVitals/WebVitalRender'
import React from 'react'

import { HighlightEvent } from '../HighlightEvent'

export type EventRenderDetails = {
	title?: string
	payload?: string
	displayValue: string | React.ReactNode
	isReactNode?: boolean
}

export const getEventRenderDetails = (
	e: HighlightEvent,
): EventRenderDetails => {
	const details: EventRenderDetails = {
		displayValue: '',
	}
	if (e.type === EventType.Custom) {
		const payload = e.data.payload as any

		details.title = e.data.tag
		switch (e.data.tag) {
			case 'Identify':
				details.displayValue = JSON.parse(payload).user_identifier
				break
			case 'Track':
				try {
					const json = JSON.parse(payload)
					details.displayValue = json.event
				} catch {
					details.displayValue = e.identifier
				}
				break
			case 'Viewport':
				details.displayValue = `${payload.height} x ${payload.width}`
				break
			case 'Navigate':
			case 'Click':
			case 'Focus':
			case 'Segment':
				try {
					const keys = Object.keys(JSON.parse(payload))
					details.displayValue = `{${keys.join(', ')}}`
				} catch {
					details.displayValue = payload
				}
				break
			case 'Web Vitals':
				details.displayValue = (
					<WebVitalSimpleRenderer vitals={payload.vitals} />
				)
				details.isReactNode = true
				break
			case 'Page Unload':
				details.displayValue = 'Page Unload'
				break
			case 'TabHidden':
				details.displayValue = payload ? 'Tab Hidden' : 'Tab Visible'
				break
			case 'RageClicks':
				details.displayValue = `Total clicks: ${payload}`
				break
			default:
				details.displayValue = payload
				break
		}
		details.payload = e.data.payload as string
	}

	return details
}
