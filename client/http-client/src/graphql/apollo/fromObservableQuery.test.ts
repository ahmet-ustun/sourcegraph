import { Observable as ZenObservable, type ObservableQuery } from '@apollo/client'
import delay from 'delay'
import { isObservable } from 'rxjs'
import sinon from 'sinon'

import { fromObservableQuery, fromObservableQueryPromise } from './fromObservableQuery'

const QUERY_RESULT = { data: {}, loading: false, networkStatus: 1 }
const UNSUBSCRIBE = sinon.spy()

function createObservableQuery() {
    return new ZenObservable(observer => {
        observer.next(QUERY_RESULT)
        observer.complete()

        return UNSUBSCRIBE
    }) as ObservableQuery
}

describe('fromObservableQuery', () => {
    it('converts `ObservableQuery` to `rxjs` observable', () => {
        const observable = fromObservableQuery(createObservableQuery())

        expect(isObservable(observable)).toBe(true)
    })

    it('subscribes to `ObservableQuery` data', done => {
        const observable = fromObservableQuery(createObservableQuery())

        observable.subscribe(data => {
            expect(data).toEqual(QUERY_RESULT)
            done()
        })
    })

    it('exposes `ObservableQuery` unsubscribe method', () => {
        const observable = fromObservableQuery(createObservableQuery())

        observable.subscribe().unsubscribe()
        sinon.assert.called(UNSUBSCRIBE)
    })
})

describe('fromObservableQueryPromise', () => {
    it('converts `Promise<ObservableQuery>` to `rxjs` observable', () => {
        const observable = fromObservableQueryPromise(Promise.resolve(createObservableQuery()))

        expect(isObservable(observable)).toBe(true)
    })

    it('subscribes to `ObservableQuery` data', done => {
        const observable = fromObservableQueryPromise(Promise.resolve(createObservableQuery()))

        observable.subscribe(data => {
            expect(data).toEqual(QUERY_RESULT)
            done()
        })
    })

    it('exposes `ObservableQuery` unsubscribe method', () => {
        const observable = fromObservableQueryPromise(Promise.resolve(createObservableQuery()))

        observable.subscribe().unsubscribe()
        sinon.assert.called(UNSUBSCRIBE)
    })

    it('unsubscribes on Promise rejection', async () => {
        const observable = fromObservableQueryPromise(Promise.reject<ObservableQuery>(new Error('oops')))

        const subscription = observable.subscribe()
        await delay(0)
        expect(subscription.closed).toBe(true)
    })
})
